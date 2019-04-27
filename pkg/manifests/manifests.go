package manifests

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	operatorv1 "github.com/openshift/api/operator/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/apiserver/pkg/storage/names"
	routev1 "github.com/openshift/api/route/v1"
)

const (
	RouterNamespaceAsset		= "assets/router/namespace.yaml"
	RouterServiceAccount		= "assets/router/service-account.yaml"
	RouterClusterRole		= "assets/router/cluster-role.yaml"
	RouterClusterRoleBinding	= "assets/router/cluster-role-binding.yaml"
	RouterDeploymentAsset		= "assets/router/deployment.yaml"
	RouterServiceInternal		= "assets/router/service-internal.yaml"
	RouterServiceCloud		= "assets/router/service-cloud.yaml"
	MetricsClusterRole		= "assets/router/metrics/cluster-role.yaml"
	MetricsClusterRoleBinding	= "assets/router/metrics/cluster-role-binding.yaml"
	MetricsRole			= "assets/router/metrics/role.yaml"
	MetricsRoleBinding		= "assets/router/metrics/role-binding.yaml"
	ServingCertSecretAnnotation	= "service.alpha.openshift.io/serving-cert-secret-name"
	OwningIngressControllerLabel	= "ingresscontroller.operator.openshift.io/owning-ingresscontroller"
)

func MustAssetReader(asset string) io.Reader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bytes.NewReader(MustAsset(asset))
}

type Factory struct{}

func RouterNamespace() *corev1.Namespace {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ns, err := NewNamespace(MustAssetReader(RouterNamespaceAsset))
	if err != nil {
		panic(err)
	}
	return ns
}
func (f *Factory) RouterServiceAccount() (*corev1.ServiceAccount, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sa, err := NewServiceAccount(MustAssetReader(RouterServiceAccount))
	if err != nil {
		return nil, err
	}
	return sa, nil
}
func (f *Factory) RouterClusterRole() (*rbacv1.ClusterRole, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cr, err := NewClusterRole(MustAssetReader(RouterClusterRole))
	if err != nil {
		return nil, err
	}
	return cr, nil
}
func (f *Factory) RouterClusterRoleBinding() (*rbacv1.ClusterRoleBinding, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	crb, err := NewClusterRoleBinding(MustAssetReader(RouterClusterRoleBinding))
	if err != nil {
		return nil, err
	}
	return crb, nil
}
func (f *Factory) RouterStatsSecret(cr *operatorv1.IngressController) (*corev1.Secret, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := &corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: fmt.Sprintf("router-stats-%s", cr.Name), Namespace: "openshift-ingress"}, Type: corev1.SecretTypeOpaque, Data: map[string][]byte{}}
	generatedUser := names.SimpleNameGenerator.GenerateName("user")
	generatedPassword := names.SimpleNameGenerator.GenerateName("pass")
	s.Data["statsUsername"] = []byte(base64.StdEncoding.EncodeToString([]byte(generatedUser)))
	s.Data["statsPassword"] = []byte(base64.StdEncoding.EncodeToString([]byte(generatedPassword)))
	return s, nil
}
func RouterDeployment() *appsv1.Deployment {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	deployment, err := NewDeployment(MustAssetReader(RouterDeploymentAsset))
	if err != nil {
		panic(err)
	}
	return deployment
}
func InternalIngressControllerService() *corev1.Service {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s, err := NewService(MustAssetReader(RouterServiceInternal))
	if err != nil {
		panic(err)
	}
	return s
}
func LoadBalancerService() *corev1.Service {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s, err := NewService(MustAssetReader(RouterServiceCloud))
	if err != nil {
		panic(err)
	}
	return s
}
func (f *Factory) MetricsClusterRole() (*rbacv1.ClusterRole, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cr, err := NewClusterRole(MustAssetReader(MetricsClusterRole))
	if err != nil {
		return nil, err
	}
	return cr, nil
}
func (f *Factory) MetricsClusterRoleBinding() (*rbacv1.ClusterRoleBinding, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	crb, err := NewClusterRoleBinding(MustAssetReader(MetricsClusterRoleBinding))
	if err != nil {
		return nil, err
	}
	return crb, nil
}
func (f *Factory) MetricsRole() (*rbacv1.Role, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r, err := NewRole(MustAssetReader(MetricsRole))
	if err != nil {
		return nil, err
	}
	return r, nil
}
func (f *Factory) MetricsRoleBinding() (*rbacv1.RoleBinding, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rb, err := NewRoleBinding(MustAssetReader(MetricsRoleBinding))
	if err != nil {
		return nil, err
	}
	return rb, nil
}
func NewServiceAccount(manifest io.Reader) (*corev1.ServiceAccount, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sa := corev1.ServiceAccount{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&sa); err != nil {
		return nil, err
	}
	return &sa, nil
}
func NewRole(manifest io.Reader) (*rbacv1.Role, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := rbacv1.Role{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&r); err != nil {
		return nil, err
	}
	return &r, nil
}
func NewRoleBinding(manifest io.Reader) (*rbacv1.RoleBinding, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rb := rbacv1.RoleBinding{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&rb); err != nil {
		return nil, err
	}
	return &rb, nil
}
func NewClusterRole(manifest io.Reader) (*rbacv1.ClusterRole, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cr := rbacv1.ClusterRole{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&cr); err != nil {
		return nil, err
	}
	return &cr, nil
}
func NewClusterRoleBinding(manifest io.Reader) (*rbacv1.ClusterRoleBinding, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	crb := rbacv1.ClusterRoleBinding{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&crb); err != nil {
		return nil, err
	}
	return &crb, nil
}
func NewService(manifest io.Reader) (*corev1.Service, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := corev1.Service{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&s); err != nil {
		return nil, err
	}
	return &s, nil
}
func NewNamespace(manifest io.Reader) (*corev1.Namespace, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ns := corev1.Namespace{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&ns); err != nil {
		return nil, err
	}
	return &ns, nil
}
func NewDeployment(manifest io.Reader) (*appsv1.Deployment, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o := appsv1.Deployment{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&o); err != nil {
		return nil, err
	}
	return &o, nil
}
func NewRoute(manifest io.Reader) (*routev1.Route, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	o := routev1.Route{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&o); err != nil {
		return nil, err
	}
	return &o, nil
}
