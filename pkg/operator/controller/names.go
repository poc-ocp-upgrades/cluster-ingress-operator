package controller

import (
	"fmt"
	operatorv1 "github.com/openshift/api/operator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	GlobalMachineSpecifiedConfigNamespace	= "openshift-config-managed"
	caCertSecretName			= "router-ca"
	caCertConfigMapName			= "router-ca"
	routerCertsGlobalSecretName		= "router-certs"
	controllerDeploymentLabel		= "ingresscontroller.operator.openshift.io/deployment-ingresscontroller"
)

func RouterDeploymentName(ci *operatorv1.IngressController) types.NamespacedName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return types.NamespacedName{Namespace: "openshift-ingress", Name: "router-" + ci.Name}
}
func RouterCASecretName(operatorNamespace string) types.NamespacedName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return types.NamespacedName{Namespace: operatorNamespace, Name: caCertSecretName}
}
func RouterCAConfigMapName() types.NamespacedName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return types.NamespacedName{Namespace: GlobalMachineSpecifiedConfigNamespace, Name: caCertConfigMapName}
}
func RouterCertsGlobalSecretName() types.NamespacedName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return types.NamespacedName{Namespace: GlobalMachineSpecifiedConfigNamespace, Name: routerCertsGlobalSecretName}
}
func RouterOperatorGeneratedDefaultCertificateSecretName(ci *operatorv1.IngressController, namespace string) types.NamespacedName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return types.NamespacedName{Namespace: namespace, Name: fmt.Sprintf("router-certs-%s", ci.Name)}
}
func RouterEffectiveDefaultCertificateSecretName(ci *operatorv1.IngressController, namespace string) types.NamespacedName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if cert := ci.Spec.DefaultCertificate; cert != nil {
		return types.NamespacedName{Namespace: namespace, Name: cert.Name}
	}
	return RouterOperatorGeneratedDefaultCertificateSecretName(ci, namespace)
}
func IngressControllerDeploymentLabel(ic *operatorv1.IngressController) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ic.Name
}
func IngressControllerDeploymentPodSelector(ic *operatorv1.IngressController) *metav1.LabelSelector {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &metav1.LabelSelector{MatchLabels: map[string]string{controllerDeploymentLabel: IngressControllerDeploymentLabel(ic)}}
}
func InternalIngressControllerServiceName(ic *operatorv1.IngressController) types.NamespacedName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return types.NamespacedName{Namespace: "openshift-ingress", Name: "router-internal-" + ic.Name}
}
