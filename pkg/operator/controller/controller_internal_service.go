package controller

import (
	"context"
	"fmt"
	operatorv1 "github.com/openshift/api/operator/v1"
	"github.com/openshift/cluster-ingress-operator/pkg/manifests"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ServingCertSecretAnnotation = "service.alpha.openshift.io/serving-cert-secret-name"
)

func (r *reconciler) ensureInternalIngressControllerService(ic *operatorv1.IngressController, deploymentRef metav1.OwnerReference) (*corev1.Service, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	desired := desiredInternalIngressControllerService(ic, deploymentRef)
	current, err := r.currentInternalIngressControllerService(ic)
	if err != nil {
		return nil, err
	}
	if current != nil {
		return current, nil
	}
	if err := r.client.Create(context.TODO(), desired); err != nil {
		return nil, fmt.Errorf("failed to create internal ingresscontroller service: %v", err)
	}
	log.Info("created internal ingresscontroller service", "service", desired)
	return desired, nil
}
func (r *reconciler) currentInternalIngressControllerService(ic *operatorv1.IngressController) (*corev1.Service, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	current := &corev1.Service{}
	err := r.client.Get(context.TODO(), InternalIngressControllerServiceName(ic), current)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return current, nil
}
func desiredInternalIngressControllerService(ic *operatorv1.IngressController, deploymentRef metav1.OwnerReference) *corev1.Service {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := manifests.InternalIngressControllerService()
	name := InternalIngressControllerServiceName(ic)
	s.Namespace = name.Namespace
	s.Name = name.Name
	s.Labels = map[string]string{manifests.OwningIngressControllerLabel: ic.Name}
	s.Annotations = map[string]string{ServingCertSecretAnnotation: fmt.Sprintf("router-metrics-certs-%s", ic.Name)}
	s.Spec.Selector = IngressControllerDeploymentPodSelector(ic).MatchLabels
	s.SetOwnerReferences([]metav1.OwnerReference{deploymentRef})
	return s
}
