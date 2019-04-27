package controller

import (
	"fmt"
	operatorv1 "github.com/openshift/api/operator/v1"
	"github.com/openshift/cluster-ingress-operator/pkg/dns"
	corev1 "k8s.io/api/core/v1"
	configv1 "github.com/openshift/api/config/v1"
)

func (r *reconciler) ensureDNS(ci *operatorv1.IngressController, service *corev1.Service, dnsConfig *configv1.DNS) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ingress := service.Status.LoadBalancer.Ingress
	if len(ingress) == 0 || len(ingress[0].Hostname) == 0 {
		return fmt.Errorf("no load balancer is assigned to service %s/%s", service.Namespace, service.Name)
	}
	dnsRecords, err := desiredDNSRecords(ci, ingress[0].Hostname, dnsConfig)
	if err != nil {
		return err
	}
	for _, record := range dnsRecords {
		err := r.DNSManager.Ensure(record)
		if err != nil {
			return fmt.Errorf("failed to ensure DNS record %v for %s/%s: %v", record, ci.Namespace, ci.Name, err)
		}
		log.Info("ensured DNS record for ingresscontroller", "namespace", ci.Namespace, "name", ci.Name, "record", record)
	}
	return nil
}
func desiredDNSRecords(ci *operatorv1.IngressController, hostname string, dnsConfig *configv1.DNS) ([]*dns.Record, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	records := []*dns.Record{}
	if len(ci.Status.Domain) == 0 {
		return records, nil
	}
	if ci.Status.EndpointPublishingStrategy.Type != operatorv1.LoadBalancerServiceStrategyType {
		return records, nil
	}
	if dnsConfig.Spec.PrivateZone == nil && dnsConfig.Spec.PublicZone == nil {
		return records, nil
	}
	domain := fmt.Sprintf("*.%s", ci.Status.Domain)
	makeRecord := func(zone *configv1.DNSZone) *dns.Record {
		return &dns.Record{Zone: *zone, Type: dns.ALIASRecord, Alias: &dns.AliasRecord{Domain: domain, Target: hostname}}
	}
	if dnsConfig.Spec.PrivateZone != nil {
		records = append(records, makeRecord(dnsConfig.Spec.PrivateZone))
	}
	if dnsConfig.Spec.PublicZone != nil {
		records = append(records, makeRecord(dnsConfig.Spec.PublicZone))
	}
	return records, nil
}
