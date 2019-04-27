package certificate

import (
	"context"
	"fmt"
	operatorv1 "github.com/openshift/api/operator/v1"
	"github.com/openshift/cluster-ingress-operator/pkg/operator/controller"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (r *reconciler) ensureRouterCAConfigMap(secret *corev1.Secret, ingresses []operatorv1.IngressController) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	desired, err := desiredRouterCAConfigMap(secret, ingresses)
	if err != nil {
		return err
	}
	current, err := r.currentRouterCAConfigMap()
	if err != nil {
		return err
	}
	switch {
	case desired == nil && current == nil:
	case desired == nil && current != nil:
		if deleted, err := r.deleteRouterCAConfigMap(current); err != nil {
			return fmt.Errorf("failed to ensure router CA was unpublished: %v", err)
		} else if deleted {
			r.recorder.Eventf(current, "Normal", "UnpublishedDefaultRouterCA", "Unpublished default router CA")
		}
	case desired != nil && current == nil:
		if created, err := r.createRouterCAConfigMap(desired); err != nil {
			return fmt.Errorf("failed to ensure router CA was published: %v", err)
		} else if created {
			new, err := r.currentRouterCAConfigMap()
			if err != nil {
				return err
			}
			r.recorder.Eventf(new, "Normal", "PublishedDefaultRouterCA", "Published default router CA")
		}
	case desired != nil && current != nil:
		if updated, err := r.updateRouterCAConfigMap(current, desired); err != nil {
			return fmt.Errorf("failed to update published router CA: %v", err)
		} else if updated {
			r.recorder.Eventf(current, "Normal", "UpdatedPublishedDefaultRouterCA", "Updated the published default router CA")
		}
	}
	return nil
}
func desiredRouterCAConfigMap(secret *corev1.Secret, ingresses []operatorv1.IngressController) (*corev1.ConfigMap, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !shouldPublishRouterCA(ingresses) {
		return nil, nil
	}
	name := controller.RouterCAConfigMapName()
	cm := &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: name.Name, Namespace: name.Namespace}, Data: map[string]string{"ca-bundle.crt": string(secret.Data["tls.crt"])}}
	return cm, nil
}
func shouldPublishRouterCA(ingresses []operatorv1.IngressController) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, ci := range ingresses {
		if ci.Spec.DefaultCertificate == nil {
			return true
		}
	}
	return false
}
func (r *reconciler) currentRouterCAConfigMap() (*corev1.ConfigMap, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	name := controller.RouterCAConfigMapName()
	cm := &corev1.ConfigMap{}
	if err := r.client.Get(context.TODO(), name, cm); err != nil {
		if errors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return cm, nil
}
func (r *reconciler) createRouterCAConfigMap(cm *corev1.ConfigMap) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := r.client.Create(context.TODO(), cm); err != nil {
		return false, err
	}
	return true, nil
}
func (r *reconciler) updateRouterCAConfigMap(current, desired *corev1.ConfigMap) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if routerCAConfigMapsEqual(current, desired) {
		return false, nil
	}
	updated := current.DeepCopy()
	updated.Data = desired.Data
	if err := r.client.Update(context.TODO(), updated); err != nil {
		return false, err
	}
	return true, nil
}
func (r *reconciler) deleteRouterCAConfigMap(cm *corev1.ConfigMap) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := r.client.Delete(context.TODO(), cm); err != nil {
		if errors.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func routerCAConfigMapsEqual(a, b *corev1.ConfigMap) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if a.Data["ca-bundle.crt"] != b.Data["ca-bundle.crt"] {
		return false
	}
	return true
}
