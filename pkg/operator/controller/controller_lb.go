package controller

import (
	"context"
	"fmt"
	operatorv1 "github.com/openshift/api/operator/v1"
	"github.com/openshift/cluster-ingress-operator/pkg/manifests"
	"github.com/openshift/cluster-ingress-operator/pkg/util/slice"
	corev1 "k8s.io/api/core/v1"
	configv1 "github.com/openshift/api/config/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

const (
	loadBalancerServiceFinalizer	= "ingress.openshift.io/operator"
	awsLBProxyProtocolAnnotation	= "service.beta.kubernetes.io/aws-load-balancer-proxy-protocol"
)

func (r *reconciler) ensureLoadBalancerService(ci *operatorv1.IngressController, deploymentRef metav1.OwnerReference, infraConfig *configv1.Infrastructure) (*corev1.Service, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	desiredLBService, err := desiredLoadBalancerService(ci, deploymentRef, infraConfig)
	if err != nil {
		return nil, err
	}
	currentLBService, err := r.currentLoadBalancerService(ci)
	if err != nil {
		return nil, err
	}
	if desiredLBService != nil && currentLBService == nil {
		if err := r.client.Create(context.TODO(), desiredLBService); err != nil {
			return nil, fmt.Errorf("failed to create load balancer service %s/%s: %v", desiredLBService.Namespace, desiredLBService.Name, err)
		}
		log.Info("created load balancer service", "namespace", desiredLBService.Namespace, "name", desiredLBService.Name)
		return desiredLBService, nil
	}
	return currentLBService, nil
}
func loadBalancerServiceName(ci *operatorv1.IngressController) types.NamespacedName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return types.NamespacedName{Namespace: "openshift-ingress", Name: "router-" + ci.Name}
}
func desiredLoadBalancerService(ci *operatorv1.IngressController, deploymentRef metav1.OwnerReference, infraConfig *configv1.Infrastructure) (*corev1.Service, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ci.Status.EndpointPublishingStrategy.Type != operatorv1.LoadBalancerServiceStrategyType {
		return nil, nil
	}
	service := manifests.LoadBalancerService()
	name := loadBalancerServiceName(ci)
	service.Namespace = name.Namespace
	service.Name = name.Name
	if service.Labels == nil {
		service.Labels = map[string]string{}
	}
	service.Labels["router"] = name.Name
	service.Labels[manifests.OwningIngressControllerLabel] = ci.Name
	service.Spec.Selector = IngressControllerDeploymentPodSelector(ci).MatchLabels
	if infraConfig.Status.Platform == configv1.AWSPlatformType {
		if service.Annotations == nil {
			service.Annotations = map[string]string{}
		}
		service.Annotations[awsLBProxyProtocolAnnotation] = "*"
	}
	service.SetOwnerReferences([]metav1.OwnerReference{deploymentRef})
	service.Finalizers = []string{loadBalancerServiceFinalizer}
	return service, nil
}
func (r *reconciler) currentLoadBalancerService(ci *operatorv1.IngressController) (*corev1.Service, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	service := &corev1.Service{}
	if err := r.client.Get(context.TODO(), loadBalancerServiceName(ci), service); err != nil {
		if errors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return service, nil
}
func (r *reconciler) finalizeLoadBalancerService(ci *operatorv1.IngressController, dnsConfig *configv1.DNS) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	service, err := r.currentLoadBalancerService(ci)
	if err != nil {
		return err
	}
	if service == nil {
		return nil
	}
	ingress := service.Status.LoadBalancer.Ingress
	if len(ingress) > 0 && len(ingress[0].Hostname) > 0 {
		records, err := desiredDNSRecords(ci, ingress[0].Hostname, dnsConfig)
		if err != nil {
			return err
		}
		dnsErrors := []error{}
		for _, record := range records {
			if err := r.DNSManager.Delete(record); err != nil {
				dnsErrors = append(dnsErrors, fmt.Errorf("failed to delete DNS record %v for ingress %s/%s: %v", record, ci.Namespace, ci.Name, err))
			} else {
				log.Info("deleted DNS record for ingress", "namespace", ci.Namespace, "name", ci.Name, "record", record)
			}
		}
		if err := utilerrors.NewAggregate(dnsErrors); err != nil {
			return err
		}
	}
	updated := service.DeepCopy()
	if slice.ContainsString(updated.Finalizers, loadBalancerServiceFinalizer) {
		updated.Finalizers = slice.RemoveString(updated.Finalizers, loadBalancerServiceFinalizer)
		if err := r.client.Update(context.TODO(), updated); err != nil {
			return fmt.Errorf("failed to remove finalizer from service %s for ingress %s/%s: %v", service.Namespace, service.Name, ci.Name, err)
		}
	}
	return nil
}
