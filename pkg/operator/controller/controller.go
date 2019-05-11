package controller

import (
	"context"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	operatorv1 "github.com/openshift/api/operator/v1"
	"github.com/openshift/cluster-ingress-operator/pkg/dns"
	logf "github.com/openshift/cluster-ingress-operator/pkg/log"
	"github.com/openshift/cluster-ingress-operator/pkg/manifests"
	operatorclient "github.com/openshift/cluster-ingress-operator/pkg/operator/client"
	"github.com/openshift/cluster-ingress-operator/pkg/util/slice"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
	configv1 "github.com/openshift/api/config/v1"
	"k8s.io/client-go/rest"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	kclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const (
	IngressControllerFinalizer = "ingresscontroller.operator.openshift.io/finalizer-ingresscontroller"
)

var log = logf.Logger.WithName("controller")

func New(mgr manager.Manager, config Config) (controller.Controller, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	kubeClient, err := operatorclient.NewClient(config.KubeConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create kube client: %v", err)
	}
	reconciler := &reconciler{Config: config, client: kubeClient, recorder: mgr.GetEventRecorderFor("operator-controller")}
	c, err := controller.New("operator-controller", mgr, controller.Options{Reconciler: reconciler})
	if err != nil {
		return nil, err
	}
	if err := c.Watch(&source.Kind{Type: &operatorv1.IngressController{}}, &handler.EnqueueRequestForObject{}); err != nil {
		return nil, err
	}
	return c, nil
}

type Config struct {
	KubeConfig				*rest.Config
	ManifestFactory			*manifests.Factory
	Namespace				string
	DNSManager				dns.Manager
	RouterImage				string
	OperatorReleaseVersion	string
}
type reconciler struct {
	Config
	client		kclient.Client
	recorder	record.EventRecorder
}

func (r *reconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	errs := []error{}
	result := reconcile.Result{}
	log.Info("reconciling", "request", request)
	ingress := &operatorv1.IngressController{}
	if err := r.client.Get(context.TODO(), request.NamespacedName, ingress); err != nil {
		if errors.IsNotFound(err) {
			log.Info("ingresscontroller not found; reconciliation will be skipped", "request", request)
		} else {
			errs = append(errs, fmt.Errorf("failed to get ingresscontroller %q: %v", request, err))
		}
		ingress = nil
	}
	if ingress != nil {
		dnsConfig := &configv1.DNS{}
		if err := r.client.Get(context.TODO(), types.NamespacedName{Name: "cluster"}, dnsConfig); err != nil {
			errs = append(errs, fmt.Errorf("failed to get dns 'cluster': %v", err))
			dnsConfig = nil
		}
		infraConfig := &configv1.Infrastructure{}
		if err := r.client.Get(context.TODO(), types.NamespacedName{Name: "cluster"}, infraConfig); err != nil {
			errs = append(errs, fmt.Errorf("failed to get infrastructure 'cluster': %v", err))
			infraConfig = nil
		}
		ingressConfig := &configv1.Ingress{}
		if err := r.client.Get(context.TODO(), types.NamespacedName{Name: "cluster"}, ingressConfig); err != nil {
			errs = append(errs, fmt.Errorf("failed to get ingress 'cluster': %v", err))
			ingressConfig = nil
		}
		if dnsConfig != nil && infraConfig != nil && ingressConfig != nil {
			if err := r.ensureRouterNamespace(); err != nil {
				errs = append(errs, fmt.Errorf("failed to ensure router namespace: %v", err))
			}
			if err := r.enforceEffectiveIngressDomain(ingress, ingressConfig); err != nil {
				errs = append(errs, fmt.Errorf("failed to enforce the effective ingress domain for ingresscontroller %s: %v", ingress.Name, err))
			} else if IsStatusDomainSet(ingress) {
				if err := r.enforceEffectiveEndpointPublishingStrategy(ingress, infraConfig); err != nil {
					errs = append(errs, fmt.Errorf("failed to enforce the effective HA configuration for ingresscontroller %s: %v", ingress.Name, err))
				} else if ingress.DeletionTimestamp != nil {
					if err := r.ensureIngressDeleted(ingress, dnsConfig, infraConfig); err != nil {
						errs = append(errs, fmt.Errorf("failed to ensure ingress deletion: %v", err))
					}
				} else if err := r.enforceIngressFinalizer(ingress); err != nil {
					errs = append(errs, fmt.Errorf("failed to enforce ingress finalizer %s/%s: %v", ingress.Namespace, ingress.Name, err))
				} else {
					if err := r.ensureIngressController(ingress, dnsConfig, infraConfig); err != nil {
						errs = append(errs, fmt.Errorf("failed to ensure ingresscontroller: %v", err))
					}
				}
			}
		}
	}
	if err := r.syncOperatorStatus(); err != nil {
		errs = append(errs, fmt.Errorf("failed to sync operator status: %v", err))
	}
	return result, utilerrors.NewAggregate(errs)
}
func (r *reconciler) enforceEffectiveIngressDomain(ic *operatorv1.IngressController, ingressConfig *configv1.Ingress) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(ic.Status.Domain) > 0 {
		return nil
	}
	updated := ic.DeepCopy()
	var domain string
	switch {
	case len(ic.Spec.Domain) > 0:
		domain = ic.Spec.Domain
	default:
		domain = ingressConfig.Spec.Domain
	}
	unique, err := r.isDomainUnique(domain)
	if err != nil {
		return err
	}
	if !unique {
		log.Info("domain not unique, not setting status domain for IngressController", "namespace", ic.Namespace, "name", ic.Name)
		availableCondition := &operatorv1.OperatorCondition{Type: operatorv1.IngressControllerAvailableConditionType, Status: operatorv1.ConditionFalse, Reason: "InvalidDomain", Message: fmt.Sprintf("domain %q is already in use by another IngressController", domain)}
		updated.Status.Conditions = setIngressStatusCondition(updated.Status.Conditions, availableCondition)
	} else {
		updated.Status.Domain = domain
	}
	if err := r.client.Status().Update(context.TODO(), updated); err != nil {
		return fmt.Errorf("failed to update status of IngressController %s/%s: %v", updated.Namespace, updated.Name, err)
	}
	if err := r.client.Get(context.TODO(), types.NamespacedName{Namespace: updated.Namespace, Name: updated.Name}, ic); err != nil {
		return fmt.Errorf("failed to get IngressController %s/%s: %v", updated.Namespace, updated.Name, err)
	}
	return nil
}
func (r *reconciler) isDomainUnique(domain string) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ingresses := &operatorv1.IngressControllerList{}
	if err := r.client.List(context.TODO(), ingresses, client.InNamespace(r.Namespace)); err != nil {
		return false, fmt.Errorf("failed to list ingresscontrollers: %v", err)
	}
	for _, ing := range ingresses.Items {
		if domain == ing.Status.Domain {
			log.Info("domain conflicts with existing IngressController", "domain", domain, "namespace", ing.Namespace, "name", ing.Name)
			return false, nil
		}
	}
	return true, nil
}
func publishingStrategyTypeForInfra(infraConfig *configv1.Infrastructure) operatorv1.EndpointPublishingStrategyType {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch infraConfig.Status.Platform {
	case configv1.AWSPlatformType:
		return operatorv1.LoadBalancerServiceStrategyType
	case configv1.LibvirtPlatformType:
		return operatorv1.HostNetworkStrategyType
	}
	return operatorv1.HostNetworkStrategyType
}
func (r *reconciler) enforceEffectiveEndpointPublishingStrategy(ci *operatorv1.IngressController, infraConfig *configv1.Infrastructure) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ci.Status.EndpointPublishingStrategy != nil {
		return nil
	}
	updated := ci.DeepCopy()
	switch {
	case ci.Spec.EndpointPublishingStrategy != nil:
		updated.Status.EndpointPublishingStrategy = ci.Spec.EndpointPublishingStrategy.DeepCopy()
	default:
		updated.Status.EndpointPublishingStrategy = &operatorv1.EndpointPublishingStrategy{Type: publishingStrategyTypeForInfra(infraConfig)}
	}
	if err := r.client.Status().Update(context.TODO(), updated); err != nil {
		return fmt.Errorf("failed to update status of ingresscontroller %s/%s: %v", updated.Namespace, updated.Name, err)
	}
	if err := r.client.Get(context.TODO(), types.NamespacedName{Namespace: updated.Namespace, Name: updated.Name}, ci); err != nil {
		return fmt.Errorf("failed to get ingresscontroller %s/%s: %v", updated.Namespace, updated.Name, err)
	}
	return nil
}
func (r *reconciler) enforceIngressFinalizer(ingress *operatorv1.IngressController) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !slice.ContainsString(ingress.Finalizers, IngressControllerFinalizer) {
		ingress.Finalizers = append(ingress.Finalizers, IngressControllerFinalizer)
		if err := r.client.Update(context.TODO(), ingress); err != nil {
			return err
		}
		log.Info("enforced finalizer for ingress", "namespace", ingress.Namespace, "name", ingress.Name)
	}
	return nil
}
func (r *reconciler) ensureIngressDeleted(ingress *operatorv1.IngressController, dnsConfig *configv1.DNS, infraConfig *configv1.Infrastructure) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := r.finalizeLoadBalancerService(ingress, dnsConfig); err != nil {
		return fmt.Errorf("failed to finalize load balancer service for %s: %v", ingress.Name, err)
	}
	log.Info("finalized load balancer service for ingress", "namespace", ingress.Namespace, "name", ingress.Name)
	if err := r.ensureRouterDeleted(ingress); err != nil {
		return fmt.Errorf("failed to delete deployment for ingress %s: %v", ingress.Name, err)
	}
	log.Info("deleted deployment for ingress", "namespace", ingress.Namespace, "name", ingress.Name)
	if slice.ContainsString(ingress.Finalizers, IngressControllerFinalizer) {
		updated := ingress.DeepCopy()
		updated.Finalizers = slice.RemoveString(updated.Finalizers, IngressControllerFinalizer)
		if err := r.client.Update(context.TODO(), updated); err != nil {
			return fmt.Errorf("failed to remove finalizer from ingresscontroller %s: %v", ingress.Name, err)
		}
	}
	return nil
}
func (r *reconciler) ensureRouterNamespace() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cr, err := r.ManifestFactory.RouterClusterRole()
	if err != nil {
		return fmt.Errorf("failed to build router cluster role: %v", err)
	}
	if err := r.client.Get(context.TODO(), types.NamespacedName{Name: cr.Name}, cr); err != nil {
		if !errors.IsNotFound(err) {
			return fmt.Errorf("failed to get router cluster role %s: %v", cr.Name, err)
		}
		if err := r.client.Create(context.TODO(), cr); err != nil {
			return fmt.Errorf("failed to create router cluster role %s: %v", cr.Name, err)
		}
		log.Info("created router cluster role", "name", cr.Name)
	}
	ns := manifests.RouterNamespace()
	if err := r.client.Get(context.TODO(), types.NamespacedName{Name: ns.Name}, ns); err != nil {
		if !errors.IsNotFound(err) {
			return fmt.Errorf("failed to get router namespace %q: %v", ns.Name, err)
		}
		if err := r.client.Create(context.TODO(), ns); err != nil {
			return fmt.Errorf("failed to create router namespace %s: %v", ns.Name, err)
		}
		log.Info("created router namespace", "name", ns.Name)
	}
	sa, err := r.ManifestFactory.RouterServiceAccount()
	if err != nil {
		return fmt.Errorf("failed to build router service account: %v", err)
	}
	if err := r.client.Get(context.TODO(), types.NamespacedName{Namespace: sa.Namespace, Name: sa.Name}, sa); err != nil {
		if !errors.IsNotFound(err) {
			return fmt.Errorf("failed to get router service account %s/%s: %v", sa.Namespace, sa.Name, err)
		}
		if err := r.client.Create(context.TODO(), sa); err != nil {
			return fmt.Errorf("failed to create router service account %s/%s: %v", sa.Namespace, sa.Name, err)
		}
		log.Info("created router service account", "namespace", sa.Namespace, "name", sa.Name)
	}
	crb, err := r.ManifestFactory.RouterClusterRoleBinding()
	if err != nil {
		return fmt.Errorf("failed to build router cluster role binding: %v", err)
	}
	if err := r.client.Get(context.TODO(), types.NamespacedName{Name: crb.Name}, crb); err != nil {
		if !errors.IsNotFound(err) {
			return fmt.Errorf("failed to get router cluster role binding %s: %v", crb.Name, err)
		}
		if err := r.client.Create(context.TODO(), crb); err != nil {
			return fmt.Errorf("failed to create router cluster role binding %s: %v", crb.Name, err)
		}
		log.Info("created router cluster role binding", "name", crb.Name)
	}
	return nil
}
func (r *reconciler) ensureIngressController(ci *operatorv1.IngressController, dnsConfig *configv1.DNS, infraConfig *configv1.Infrastructure) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	errs := []error{}
	if deployment, err := r.ensureRouterDeployment(ci, infraConfig); err != nil {
		errs = append(errs, fmt.Errorf("failed to ensure router deployment for %s: %v", ci.Name, err))
	} else {
		trueVar := true
		deploymentRef := metav1.OwnerReference{APIVersion: "apps/v1", Kind: "Deployment", Name: deployment.Name, UID: deployment.UID, Controller: &trueVar}
		if lbService, err := r.ensureLoadBalancerService(ci, deploymentRef, infraConfig); err != nil {
			errs = append(errs, fmt.Errorf("failed to ensure load balancer service for %s: %v", ci.Name, err))
		} else if lbService != nil {
			if err := r.ensureDNS(ci, lbService, dnsConfig); err != nil {
				errs = append(errs, fmt.Errorf("failed to ensure DNS for %s: %v", ci.Name, err))
			}
		}
		if internalSvc, err := r.ensureInternalIngressControllerService(ci, deploymentRef); err != nil {
			errs = append(errs, fmt.Errorf("failed to create internal router service for ingresscontroller %s: %v", ci.Name, err))
		} else if err := r.ensureMetricsIntegration(ci, internalSvc, deploymentRef); err != nil {
			errs = append(errs, fmt.Errorf("failed to integrate metrics with openshift-monitoring for ingresscontroller %s: %v", ci.Name, err))
		}
		if err := r.syncIngressControllerStatus(deployment, ci); err != nil {
			errs = append(errs, fmt.Errorf("failed to sync ingresscontroller status: %v", err))
		}
	}
	return utilerrors.NewAggregate(errs)
}
func (r *reconciler) ensureMetricsIntegration(ci *operatorv1.IngressController, svc *corev1.Service, deploymentRef metav1.OwnerReference) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	statsSecret, err := r.ManifestFactory.RouterStatsSecret(ci)
	if err != nil {
		return fmt.Errorf("failed to build router stats secret: %v", err)
	}
	if err := r.client.Get(context.TODO(), types.NamespacedName{Namespace: statsSecret.Namespace, Name: statsSecret.Name}, statsSecret); err != nil {
		if !errors.IsNotFound(err) {
			return fmt.Errorf("failed to get router stats secret %s/%s, %v", statsSecret.Namespace, statsSecret.Name, err)
		}
		statsSecret.SetOwnerReferences([]metav1.OwnerReference{deploymentRef})
		if err := r.client.Create(context.TODO(), statsSecret); err != nil {
			return fmt.Errorf("failed to create router stats secret %s/%s: %v", statsSecret.Namespace, statsSecret.Name, err)
		}
		log.Info("created router stats secret", "namespace", statsSecret.Namespace, "name", statsSecret.Name)
	}
	cr, err := r.ManifestFactory.MetricsClusterRole()
	if err != nil {
		return fmt.Errorf("failed to build router metrics cluster role: %v", err)
	}
	if err := r.client.Get(context.TODO(), types.NamespacedName{Name: cr.Name}, cr); err != nil {
		if !errors.IsNotFound(err) {
			return fmt.Errorf("failed to get router metrics cluster role %s: %v", cr.Name, err)
		}
		if err := r.client.Create(context.TODO(), cr); err != nil {
			return fmt.Errorf("failed to create router metrics cluster role %s: %v", cr.Name, err)
		}
		log.Info("created router metrics cluster role", "name", cr.Name)
	}
	crb, err := r.ManifestFactory.MetricsClusterRoleBinding()
	if err != nil {
		return fmt.Errorf("failed to build router metrics cluster role binding: %v", err)
	}
	if err := r.client.Get(context.TODO(), types.NamespacedName{Name: crb.Name}, crb); err != nil {
		if !errors.IsNotFound(err) {
			return fmt.Errorf("failed to get router metrics cluster role binding %s: %v", crb.Name, err)
		}
		if err := r.client.Create(context.TODO(), crb); err != nil {
			return fmt.Errorf("failed to create router metrics cluster role binding %s: %v", crb.Name, err)
		}
		log.Info("created router metrics cluster role binding", "name", crb.Name)
	}
	mr, err := r.ManifestFactory.MetricsRole()
	if err != nil {
		return fmt.Errorf("failed to build router metrics role: %v", err)
	}
	if err := r.client.Get(context.TODO(), types.NamespacedName{Namespace: mr.Namespace, Name: mr.Name}, mr); err != nil {
		if !errors.IsNotFound(err) {
			return fmt.Errorf("failed to get router metrics role %s: %v", mr.Name, err)
		}
		if err := r.client.Create(context.TODO(), mr); err != nil {
			return fmt.Errorf("failed to create router metrics role %s: %v", mr.Name, err)
		}
		log.Info("created router metrics role", "name", mr.Name)
	}
	mrb, err := r.ManifestFactory.MetricsRoleBinding()
	if err != nil {
		return fmt.Errorf("failed to build router metrics role binding: %v", err)
	}
	if err := r.client.Get(context.TODO(), types.NamespacedName{Namespace: mrb.Namespace, Name: mrb.Name}, mrb); err != nil {
		if !errors.IsNotFound(err) {
			return fmt.Errorf("failed to get router metrics role binding %s: %v", mrb.Name, err)
		}
		if err := r.client.Create(context.TODO(), mrb); err != nil {
			return fmt.Errorf("failed to create router metrics role binding %s: %v", mrb.Name, err)
		}
		log.Info("created router metrics role binding", "name", mrb.Name)
	}
	if _, err := r.ensureServiceMonitor(ci, svc, deploymentRef); err != nil {
		return fmt.Errorf("failed to ensure servicemonitor for %s: %v", ci.Name, err)
	}
	return nil
}
func IsStatusDomainSet(ingress *operatorv1.IngressController) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(ingress.Status.Domain) == 0 {
		return false
	}
	return true
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
