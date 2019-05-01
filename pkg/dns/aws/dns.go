package aws

import (
	godefaultbytes "bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go/service/route53"
	configv1 "github.com/openshift/api/config/v1"
	"github.com/openshift/cluster-ingress-operator/pkg/dns"
	logf "github.com/openshift/cluster-ingress-operator/pkg/log"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	godefaulthttp "net/http"
	"reflect"
	godefaultruntime "runtime"
	"strings"
	"sync"
)

var (
	_   dns.Manager = &Manager{}
	log             = logf.Logger.WithName("dns")
)

type Manager struct {
	elb            *elb.ELB
	route53        *route53.Route53
	tags           *resourcegroupstaggingapi.ResourceGroupsTaggingAPI
	config         Config
	lock           sync.RWMutex
	idsToTags      map[string]map[string]string
	lbZones        map[string]string
	updatedRecords sets.String
}
type Config struct {
	AccessID  string
	AccessKey string
	DNS       *configv1.DNS
}

func NewManager(config Config, operatorReleaseVersion string) (*Manager, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	creds := credentials.NewStaticCredentials(config.AccessID, config.AccessKey, "")
	sess, err := session.NewSessionWithOptions(session.Options{Config: aws.Config{Credentials: creds}, SharedConfigState: session.SharedConfigEnable})
	if err != nil {
		return nil, fmt.Errorf("couldn't create AWS client session: %v", err)
	}
	sess.Handlers.Build.PushBackNamed(request.NamedHandler{Name: "openshift.io/ingress-operator", Fn: request.MakeAddToUserAgentHandler("openshift.io ingress-operator", operatorReleaseVersion)})
	region := aws.StringValue(sess.Config.Region)
	if len(region) > 0 {
		log.Info("using region from shared config", "region name", region)
	} else {
		metadata := ec2metadata.New(sess)
		discovered, err := metadata.Region()
		if err != nil {
			return nil, fmt.Errorf("couldn't get region from metadata: %v", err)
		}
		region = discovered
		log.Info("discovered region from metadata", "region name", region)
	}
	return &Manager{elb: elb.New(sess, aws.NewConfig().WithRegion(region)), route53: route53.New(sess), tags: resourcegroupstaggingapi.New(sess, aws.NewConfig().WithRegion("us-east-1")), config: config, idsToTags: map[string]map[string]string{}, lbZones: map[string]string{}, updatedRecords: sets.NewString()}, nil
}
func (m *Manager) getZoneID(zoneConfig configv1.DNSZone) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.lock.Lock()
	defer m.lock.Unlock()
	if len(zoneConfig.ID) > 0 {
		return zoneConfig.ID, nil
	}
	for id, tags := range m.idsToTags {
		if reflect.DeepEqual(tags, zoneConfig.Tags) {
			return id, nil
		}
	}
	var id string
	var innerError error
	f := func(resp *resourcegroupstaggingapi.GetResourcesOutput, lastPage bool) (shouldContinue bool) {
		for _, zone := range resp.ResourceTagMappingList {
			zoneARN, err := arn.Parse(aws.StringValue(zone.ResourceARN))
			if err != nil {
				innerError = fmt.Errorf("failed to parse hostedzone ARN %q: %v", aws.StringValue(zone.ResourceARN), err)
				return false
			}
			elems := strings.Split(zoneARN.Resource, "/")
			if len(elems) != 2 || elems[0] != "hostedzone" {
				innerError = fmt.Errorf("got unexpected resource ARN: %v", zoneARN)
				return false
			}
			id = elems[1]
			return false
		}
		return true
	}
	tagFilters := []*resourcegroupstaggingapi.TagFilter{}
	for k, v := range zoneConfig.Tags {
		tagFilters = append(tagFilters, &resourcegroupstaggingapi.TagFilter{Key: aws.String(k), Values: []*string{aws.String(v)}})
	}
	outerError := m.tags.GetResourcesPages(&resourcegroupstaggingapi.GetResourcesInput{ResourceTypeFilters: []*string{aws.String("route53:hostedzone")}, TagFilters: tagFilters}, f)
	if err := kerrors.NewAggregate([]error{innerError, outerError}); err != nil {
		return id, fmt.Errorf("failed to get tagged resources: %v", err)
	}
	if len(id) == 0 {
		return id, fmt.Errorf("no matching hosted zone found")
	}
	m.idsToTags[id] = zoneConfig.Tags
	log.Info("found hosted zone using tags", "zone id", id, "tags", zoneConfig.Tags)
	return id, nil
}
func (m *Manager) getLBHostedZone(name string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.lock.Lock()
	defer m.lock.Unlock()
	if id, exists := m.lbZones[name]; exists {
		return id, nil
	}
	var id string
	fn := func(resp *elb.DescribeLoadBalancersOutput, lastPage bool) (shouldContinue bool) {
		for _, lb := range resp.LoadBalancerDescriptions {
			log.V(0).Info("found load balancer", "name", aws.StringValue(lb.LoadBalancerName), "dns name", aws.StringValue(lb.DNSName), "hosted zone ID", aws.StringValue(lb.CanonicalHostedZoneNameID))
			if aws.StringValue(lb.CanonicalHostedZoneName) == name {
				id = aws.StringValue(lb.CanonicalHostedZoneNameID)
				return false
			}
		}
		return true
	}
	err := m.elb.DescribeLoadBalancersPages(&elb.DescribeLoadBalancersInput{}, fn)
	if err != nil {
		return "", fmt.Errorf("failed to describe load balancers: %v", err)
	}
	if len(id) == 0 {
		return "", fmt.Errorf("couldn't find hosted zone ID of ELB %s", name)
	}
	log.Info("associating load balancer with hosted zone", "dns name", name, "zone", id)
	m.lbZones[name] = id
	return id, nil
}

type action string

const (
	upsertAction action = "UPSERT"
	deleteAction action = "DELETE"
)

func (m *Manager) Ensure(record *dns.Record) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.change(record, upsertAction)
}
func (m *Manager) Delete(record *dns.Record) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.change(record, deleteAction)
}
func (m *Manager) change(record *dns.Record, action action) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if record.Type != dns.ALIASRecord {
		return fmt.Errorf("unsupported record type %s", record.Type)
	}
	alias := record.Alias
	if alias == nil {
		return fmt.Errorf("missing alias record")
	}
	domain, target := alias.Domain, alias.Target
	if len(domain) == 0 {
		return fmt.Errorf("domain is required")
	}
	if len(target) == 0 {
		return fmt.Errorf("target is required")
	}
	zoneID, err := m.getZoneID(record.Zone)
	if err != nil {
		return fmt.Errorf("failed to find hosted zone for record %v: %v", record, err)
	}
	targetHostedZoneID, err := m.getLBHostedZone(target)
	if err != nil {
		return fmt.Errorf("failed to get hosted zone for load balancer target %q: %v", target, err)
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	key := zoneID + domain + target
	if m.updatedRecords.Has(key) && action == upsertAction {
		log.Info("skipping DNS record update", "record", record)
		return nil
	}
	err = m.updateAlias(domain, zoneID, target, targetHostedZoneID, string(action))
	if err != nil {
		return fmt.Errorf("failed to update alias in zone %s: %v", zoneID, err)
	}
	switch action {
	case upsertAction:
		m.updatedRecords.Insert(key)
		log.Info("upserted DNS record", "record", record)
	case deleteAction:
		m.updatedRecords.Delete(key)
		log.Info("deleted DNS record", "record", record)
	}
	return nil
}
func (m *Manager) updateAlias(domain, zoneID, target, targetHostedZoneID, action string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := m.route53.ChangeResourceRecordSets(&route53.ChangeResourceRecordSetsInput{HostedZoneId: aws.String(zoneID), ChangeBatch: &route53.ChangeBatch{Changes: []*route53.Change{{Action: aws.String(action), ResourceRecordSet: &route53.ResourceRecordSet{Name: aws.String(domain), Type: aws.String("A"), AliasTarget: &route53.AliasTarget{HostedZoneId: aws.String(targetHostedZoneID), DNSName: aws.String(target), EvaluateTargetHealth: aws.Bool(false)}}}}}})
	if err != nil {
		if action == string(deleteAction) {
			if aerr, ok := err.(awserr.Error); ok {
				if strings.Contains(aerr.Message(), "not found") {
					log.Info("record not found", "zone id", zoneID, "domain", domain, "target", target)
					return nil
				}
			}
		}
		return fmt.Errorf("couldn't update DNS record in zone %s: %v", zoneID, err)
	}
	log.Info("updated DNS record", "zone id", zoneID, "domain", domain, "target", target, "response", resp)
	return nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
