package certificatepublisher

import (
	"bytes"
	"context"
	"fmt"
	operatorv1 "github.com/openshift/api/operator/v1"
	"github.com/openshift/cluster-ingress-operator/pkg/operator/controller"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"reflect"
)

func (r *reconciler) ensureRouterCertsGlobalSecret(secrets []corev1.Secret, ingresses []operatorv1.IngressController) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	desired, err := desiredRouterCertsGlobalSecret(secrets, ingresses, r.operandNamespace)
	if err != nil {
		return err
	}
	current, err := r.currentRouterCertsGlobalSecret()
	if err != nil {
		return err
	}
	switch {
	case desired == nil && current == nil:
	case desired == nil && current != nil:
		if deleted, err := r.deleteRouterCertsGlobalSecret(current); err != nil {
			return fmt.Errorf("failed to ensure router certificates secret was unpublished: %v", err)
		} else if deleted {
			r.recorder.Eventf(current, "Normal", "UnpublishedRouterCertificates", "Unpublished router certificates")
		}
	case desired != nil && current == nil:
		if created, err := r.createRouterCertsGlobalSecret(desired); err != nil {
			return fmt.Errorf("failed to ensure router certificates secret was published: %v", err)
		} else if created {
			new, err := r.currentRouterCertsGlobalSecret()
			if err != nil {
				return err
			}
			r.recorder.Eventf(new, "Normal", "PublishedRouterCertificates", "Published router certificates")
		}
	case desired != nil && current != nil:
		if updated, err := r.updateRouterCertsGlobalSecret(current, desired); err != nil {
			return fmt.Errorf("failed to update published router certificates secret: %v", err)
		} else if updated {
			r.recorder.Eventf(current, "Normal", "UpdatedPublishedRouterCertificates", "Updated the published router certificates")
		}
	}
	return nil
}
func desiredRouterCertsGlobalSecret(secrets []corev1.Secret, ingresses []operatorv1.IngressController, operandNamespace string) (*corev1.Secret, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(ingresses) == 0 || len(secrets) == 0 {
		return nil, nil
	}
	nameToSecret := map[string]*corev1.Secret{}
	for i, certSecret := range secrets {
		nameToSecret[certSecret.Name] = &secrets[i]
	}
	ingressToSecret := map[*operatorv1.IngressController]*corev1.Secret{}
	for i, ingress := range ingresses {
		name := controller.RouterEffectiveDefaultCertificateSecretName(&ingress, operandNamespace)
		if secret, ok := nameToSecret[name.Name]; ok {
			ingressToSecret[&ingresses[i]] = secret
		}
	}
	name := controller.RouterCertsGlobalSecretName()
	globalSecret := &corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: name.Name, Namespace: name.Namespace}, Data: map[string][]byte{}}
	for ingress, certSecret := range ingressToSecret {
		if len(ingress.Status.Domain) == 0 {
			continue
		}
		pem := bytes.Join([][]byte{certSecret.Data["tls.crt"], certSecret.Data["tls.key"]}, nil)
		globalSecret.Data[ingress.Status.Domain] = pem
	}
	return globalSecret, nil
}
func (r *reconciler) currentRouterCertsGlobalSecret() (*corev1.Secret, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	name := controller.RouterCertsGlobalSecretName()
	secret := &corev1.Secret{}
	if err := r.client.Get(context.TODO(), name, secret); err != nil {
		if errors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return secret, nil
}
func (r *reconciler) createRouterCertsGlobalSecret(secret *corev1.Secret) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := r.client.Create(context.TODO(), secret); err != nil {
		if errors.IsAlreadyExists(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func (r *reconciler) updateRouterCertsGlobalSecret(current, desired *corev1.Secret) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if routerCertsSecretsEqual(current, desired) {
		return false, nil
	}
	updated := current.DeepCopy()
	updated.Data = desired.Data
	if err := r.client.Update(context.TODO(), updated); err != nil {
		if errors.IsAlreadyExists(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func (r *reconciler) deleteRouterCertsGlobalSecret(secret *corev1.Secret) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := r.client.Delete(context.TODO(), secret); err != nil {
		if errors.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func routerCertsSecretsEqual(a, b *corev1.Secret) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !reflect.DeepEqual(a.Data, b.Data) {
		return false
	}
	return true
}
