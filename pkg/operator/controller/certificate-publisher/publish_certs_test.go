package certificatepublisher

import (
	"bytes"
	"testing"
	operatorv1 "github.com/openshift/api/operator/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func newSecret(name string) corev1.Secret {
	_logClusterCodePath()
	defer _logClusterCodePath()
	const (
		defaultCert	= `-----BEGIN CERTIFICATE-----
MIIDIjCCAgqgAwIBAgIBBjANBgkqhkiG9w0BAQUFADCBoTELMAkGA1UEBhMCVVMx
CzAJBgNVBAgMAlNDMRUwEwYDVQQHDAxEZWZhdWx0IENpdHkxHDAaBgNVBAoME0Rl
ZmF1bHQgQ29tcGFueSBMdGQxEDAOBgNVBAsMB1Rlc3QgQ0ExGjAYBgNVBAMMEXd3
dy5leGFtcGxlY2EuY29tMSIwIAYJKoZIhvcNAQkBFhNleGFtcGxlQGV4YW1wbGUu
Y29tMB4XDTE2MDExMzE5NDA1N1oXDTI2MDExMDE5NDA1N1owfDEYMBYGA1UEAxMP
d3d3LmV4YW1wbGUuY29tMQswCQYDVQQIEwJTQzELMAkGA1UEBhMCVVMxIjAgBgkq
hkiG9w0BCQEWE2V4YW1wbGVAZXhhbXBsZS5jb20xEDAOBgNVBAoTB0V4YW1wbGUx
EDAOBgNVBAsTB0V4YW1wbGUwgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBAM0B
u++oHV1wcphWRbMLUft8fD7nPG95xs7UeLPphFZuShIhhdAQMpvcsFeg+Bg9PWCu
v3jZljmk06MLvuWLfwjYfo9q/V+qOZVfTVHHbaIO5RTXJMC2Nn+ACF0kHBmNcbth
OOgF8L854a/P8tjm1iPR++vHnkex0NH7lyosVc/vAgMBAAGjDTALMAkGA1UdEwQC
MAAwDQYJKoZIhvcNAQEFBQADggEBADjFm5AlNH3DNT1Uzx3m66fFjqqrHEs25geT
yA3rvBuynflEHQO95M/8wCxYVyuAx4Z1i4YDC7tx0vmOn/2GXZHY9MAj1I8KCnwt
Jik7E2r1/yY0MrkawljOAxisXs821kJ+Z/51Ud2t5uhGxS6hJypbGspMS7OtBbw7
8oThK7cWtCXOldNF6ruqY1agWnhRdAq5qSMnuBXuicOP0Kbtx51a1ugE3SnvQenJ
nZxdtYUXvEsHZC/6bAtTfNh+/SwgxQJuL2ZM+VG3X2JIKY8xTDui+il7uTh422lq
wED8uwKl+bOj6xFDyw4gWoBxRobsbFaME8pkykP1+GnKDberyAM=
-----END CERTIFICATE-----
`
		defaultKey	= `-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQDNAbvvqB1dcHKYVkWzC1H7fHw+5zxvecbO1Hiz6YRWbkoSIYXQ
EDKb3LBXoPgYPT1grr942ZY5pNOjC77li38I2H6Pav1fqjmVX01Rx22iDuUU1yTA
tjZ/gAhdJBwZjXG7YTjoBfC/OeGvz/LY5tYj0fvrx55HsdDR+5cqLFXP7wIDAQAB
AoGAfE7P4Zsj6zOzGPI/Izj7Bi5OvGnEeKfzyBiH9Dflue74VRQkqqwXs/DWsNv3
c+M2Y3iyu5ncgKmUduo5X8D9To2ymPRLGuCdfZTxnBMpIDKSJ0FTwVPkr6cYyyBk
5VCbc470pQPxTAAtl2eaO1sIrzR4PcgwqrSOjwBQQocsGAECQQD8QOra/mZmxPbt
bRh8U5lhgZmirImk5RY3QMPI/1/f4k+fyjkU5FRq/yqSyin75aSAXg8IupAFRgyZ
W7BT6zwBAkEA0A0ugAGorpCbuTa25SsIOMxkEzCiKYvh0O+GfGkzWG4lkSeJqGME
keuJGlXrZNKNoCYLluAKLPmnd72X2yTL7wJARM0kAXUP0wn324w8+HQIyqqBj/gF
Vt9Q7uMQQ3s72CGu3ANZDFS2nbRZFU5koxrggk6lRRk1fOq9NvrmHg10AQJABOea
pgfj+yGLmkUw8JwgGH6xCUbHO+WBUFSlPf+Y50fJeO+OrjqPXAVKeSV3ZCwWjKT4
9viXJNJJ4WfF0bO/XwJAOMB1wQnEOSZ4v+laMwNtMq6hre5K8woqteXICoGcIWe8
u3YLAbyW/lHhOCiZu2iAI8AbmXem9lW6Tr7p/97s0w==
-----END RSA PRIVATE KEY-----
`
	)
	return corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: name}, Data: map[string][]byte{"tls.crt": []byte(defaultCert), "tls.key": []byte(defaultKey)}}
}
func newIngressController(name, defaultCertificateSecretName, domain string) operatorv1.IngressController {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ingresscontroller := operatorv1.IngressController{ObjectMeta: metav1.ObjectMeta{Name: name}, Status: operatorv1.IngressControllerStatus{Domain: domain}}
	if len(defaultCertificateSecretName) != 0 {
		ingresscontroller.Spec.DefaultCertificate = &corev1.LocalObjectReference{Name: defaultCertificateSecretName}
	}
	return ingresscontroller
}
func TestDesiredRouterCertsGlobalSecret(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	type testInputs struct {
		ingresses	[]operatorv1.IngressController
		secrets		[]corev1.Secret
	}
	type testOutputs struct{ secret *corev1.Secret }
	var (
		defaultCert	= newSecret("router-certs-default")
		defaultCI	= newIngressController("default", "", "apps.my.devcluster.openshift.com")
		ci1		= newIngressController("ci1", "s1", "dom1")
		ci2		= newIngressController("ci2", "s2", "dom2")
		s1		= newSecret("s1")
		s2		= newSecret("s2")
		data		= bytes.Join([][]byte{s1.Data["tls.crt"], s1.Data["tls.key"]}, nil)
	)
	testCases := []struct {
		description	string
		inputs		testInputs
		output		testOutputs
	}{{description: "default configuration", inputs: testInputs{[]operatorv1.IngressController{defaultCI}, []corev1.Secret{defaultCert}}, output: testOutputs{&corev1.Secret{Data: map[string][]byte{"apps.my.devcluster.openshift.com": data}}}}, {description: "no ingresses", inputs: testInputs{[]operatorv1.IngressController{}, []corev1.Secret{}}, output: testOutputs{nil}}, {description: "no secrets", inputs: testInputs{[]operatorv1.IngressController{ci1}, []corev1.Secret{}}, output: testOutputs{nil}}, {description: "missing secret", inputs: testInputs{[]operatorv1.IngressController{ci1, ci2}, []corev1.Secret{s1}}, output: testOutputs{&corev1.Secret{Data: map[string][]byte{"dom1": data}}}}, {description: "extra secret", inputs: testInputs{[]operatorv1.IngressController{ci2}, []corev1.Secret{s1, s2}}, output: testOutputs{&corev1.Secret{Data: map[string][]byte{"dom2": data}}}}, {description: "perfect match", inputs: testInputs{[]operatorv1.IngressController{ci1, ci2}, []corev1.Secret{s1, s2}}, output: testOutputs{&corev1.Secret{Data: map[string][]byte{"dom1": data, "dom2": data}}}}}
	for _, tc := range testCases {
		expected := tc.output.secret
		actual, err := desiredRouterCertsGlobalSecret(tc.inputs.secrets, tc.inputs.ingresses, "openshift-ingress")
		if err != nil {
			t.Errorf("failed to get desired router-ca global secret: %v", err)
			continue
		}
		if expected == nil || actual == nil {
			if expected != nil {
				t.Errorf("%q: expected %v, got nil", tc.description, expected)
			}
			if actual != nil {
				t.Errorf("%q: expected nil, got %v", tc.description, actual)
			}
			continue
		}
		if !routerCertsSecretsEqual(expected, actual) {
			t.Errorf("%q: expected %v, got %v", tc.description, expected, actual)
		}
	}
}
