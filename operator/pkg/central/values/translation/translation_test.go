package translation

import (
	"context"
	"testing"

	"github.com/stackrox/rox/operator/api/central/v1alpha1"
	common "github.com/stackrox/rox/operator/api/common/v1alpha1"
	"github.com/stackrox/rox/operator/pkg/values/translation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"helm.sh/helm/v3/pkg/chartutil"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
)

func TestReadBaseValues(t *testing.T) {
	_, err := chartutil.ReadValues(baseValuesYAML)
	assert.NoError(t, err)
}

func TestTranslate(t *testing.T) {
	type args struct {
		clientSet kubernetes.Interface
		c         v1alpha1.Central
	}

	connectivityPolicy := v1alpha1.ConnectivityOffline
	telemetryPolicy := v1alpha1.TelemetryEnabled
	hostPath := "/central/host/path"
	claimName := "central-claim-name"
	createClaimPolicy := v1alpha1.ClaimCreate
	scannerComponentPolicy := v1alpha1.ScannerComponentEnabled
	scannerAutoScalingPolicy := v1alpha1.ScannerAutoScalingEnabled

	truth := true
	lbPort := int32(12345)
	lbIP := "1.1.1.1"
	nodePortPort := int32(23456)
	scannerReplicas := int32(7)
	scannerMinReplicas := int32(6)
	scannerMaxReplicas := int32(8)

	tests := map[string]struct {
		args args
		want chartutil.Values
	}{
		"empty": {
			args: args{
				c: v1alpha1.Central{},
			},
			want: chartutil.Values{},
		},

		"empty spec": {
			args: args{
				c: v1alpha1.Central{
					Spec: v1alpha1.CentralSpec{},
				},
			},
			want: chartutil.Values{},
		},

		"everything and the kitchen sink": {
			args: args{
				c: v1alpha1.Central{
					ObjectMeta: v1.ObjectMeta{Namespace: "stackrox"},
					Spec: v1alpha1.CentralSpec{
						ImagePullSecrets: []corev1.LocalObjectReference{
							{Name: "image-pull-secrets-secret1"},
							{Name: "image-pull-secrets-secret2"},
						},
						Egress: &v1alpha1.Egress{
							ConnectivityPolicy: &connectivityPolicy,
							ProxyConfigSecret:  &corev1.LocalObjectReference{Name: "proxy-config-secret"},
						},
						TLS: &common.TLSConfig{
							AdditionalCAs: []common.AdditionalCA{
								{Name: "ca1-name", Content: "ca1-content"},
								{Name: "ca2-name", Content: "ca2-content"},
							},
						},
						Central: &v1alpha1.CentralComponentSpec{
							DeploymentSpec: common.DeploymentSpec{
								ServiceTLSSpec: common.ServiceTLSSpec{
									ServiceTLS: &corev1.LocalObjectReference{Name: "central-tls-spec-secret"},
								},
								NodeSelector: map[string]string{
									"central-node-selector-label1": "central-node-selector-value1",
									"central-node-selector-label2": "central-node-selector-value2",
								},
								Resources: &common.Resources{
									Override: &corev1.ResourceRequirements{
										Limits: corev1.ResourceList{
											corev1.ResourceCPU:    resource.MustParse("10"),
											corev1.ResourceMemory: resource.MustParse("20"),
										},
										Requests: corev1.ResourceList{
											corev1.ResourceCPU:    resource.MustParse("30"),
											corev1.ResourceMemory: resource.MustParse("40"),
										},
									},
								},
							},
							TelemetryPolicy:     &telemetryPolicy,
							AdminPasswordSecret: &corev1.LocalObjectReference{Name: "admin-password-secret"},
							Persistence: &v1alpha1.Persistence{
								HostPath: &hostPath,
								PersistentVolumeClaim: &v1alpha1.PersistentVolumeClaim{
									ClaimName:   &claimName,
									CreateClaim: &createClaimPolicy,
								},
							},
							Exposure: &v1alpha1.Exposure{
								LoadBalancer: &v1alpha1.ExposureLoadBalancer{
									Enabled: &truth,
									Port:    &lbPort,
									IP:      &lbIP,
								},
								NodePort: &v1alpha1.ExposureNodePort{
									Enabled: &truth,
									Port:    &nodePortPort,
								},
								Route: &v1alpha1.ExposureRoute{
									Enabled: &truth,
								},
							},
						},
						Analyzer: &v1alpha1.AnalyzerComponentSpec{
							ScannerComponent: &scannerComponentPolicy,
							Replicas: &v1alpha1.ScannerReplicas{
								AutoScaling: &scannerAutoScalingPolicy,
								Replicas:    &scannerReplicas,
								MinReplicas: &scannerMinReplicas,
								MaxReplicas: &scannerMaxReplicas,
							},
							Scanner: &common.DeploymentSpec{
								ServiceTLSSpec: common.ServiceTLSSpec{
									ServiceTLS: &corev1.LocalObjectReference{Name: "scanner-tls-spec-secret"},
								},
								NodeSelector: map[string]string{
									"scanner-node-selector-label1": "scanner-node-selector-value1",
									"scanner-node-selector-label2": "scanner-node-selector-value2",
								},
								Resources: &common.Resources{
									Override: &corev1.ResourceRequirements{
										Limits: corev1.ResourceList{
											corev1.ResourceCPU:    resource.MustParse("50"),
											corev1.ResourceMemory: resource.MustParse("60"),
										},
										Requests: corev1.ResourceList{
											corev1.ResourceCPU:    resource.MustParse("70"),
											corev1.ResourceMemory: resource.MustParse("80"),
										},
									},
								},
							},
							ScannerDB: &common.DeploymentSpec{
								ServiceTLSSpec: common.ServiceTLSSpec{
									ServiceTLS: &corev1.LocalObjectReference{Name: "scanner-db-tls-spec-secret"},
								},
								NodeSelector: map[string]string{
									"scanner-db-node-selector-label1": "scanner-db-node-selector-value1",
									"scanner-db-node-selector-label2": "scanner-db-node-selector-value2",
								},
								Resources: &common.Resources{
									Override: &corev1.ResourceRequirements{
										Limits: corev1.ResourceList{
											corev1.ResourceCPU:    resource.MustParse("90"),
											corev1.ResourceMemory: resource.MustParse("100"),
										},
										Requests: corev1.ResourceList{
											corev1.ResourceCPU:    resource.MustParse("110"),
											corev1.ResourceMemory: resource.MustParse("120"),
										},
									},
								},
							},
						},
						Customize: &common.CustomizeSpec{
							Labels: map[string]string{
								"customize-label1": "customize-label1-value",
								"customize-label2": "customize-label2-value",
							},
							Annotations: map[string]string{
								"customize-annotation1": "customize-annotation1-value",
								"customize-annotation2": "customize-annotation2-value",
							},
							EnvVars: map[string]string{
								"customize-env-var1": "customize-env-var1-value",
								"customize-env-var2": "customize-env-var2-value",
							},
						},
					},
				},
				clientSet: fake.NewSimpleClientset(
					makeSecret("proxy-config-secret", map[string]string{"config.yaml": "proxy-config-secret-content"}),
					makeSecret("central-tls-spec-secret",
						map[string]string{
							"key":  "central-tls-spec-secret-key-content",
							"cert": "central-tls-spec-secret-cert-content",
						}),
					makeSecret("admin-password-secret", map[string]string{"value": "admin-password-plaintext"}),
					makeSecret("scanner-tls-spec-secret",
						map[string]string{
							"key":  "scanner-tls-spec-secret-key-content",
							"cert": "scanner-tls-spec-secret-cert-content",
						}),
					makeSecret("scanner-db-tls-spec-secret",
						map[string]string{
							"key":  "scanner-db-tls-spec-secret-key-content",
							"cert": "scanner-db-tls-spec-secret-cert-content",
						}),
				),
			},
			want: chartutil.Values{
				"additionalCAs": map[string]interface{}{
					"ca1-name": "ca1-content",
					"ca2-name": "ca2-content",
				},
				"central": map[string]interface{}{
					"adminPassword": map[string]interface{}{
						"value": "admin-password-plaintext",
					},
					"disableTelemetry": false,
					"exposure": map[string]interface{}{
						"loadBalancer": map[string]interface{}{
							"enabled": true,
							"ip":      "1.1.1.1",
							"port":    int32(12345),
						},
						"nodePort": map[string]interface{}{
							"enabled": true,
							"port":    int32(23456),
						},
						"route": map[string]interface{}{
							"enabled": true,
						},
					},
					"nodeSelector": map[string]string{
						"central-node-selector-label1": "central-node-selector-value1",
						"central-node-selector-label2": "central-node-selector-value2",
					},
					"persistence": map[string]interface{}{
						"hostPath": "/central/host/path",
						"persistentVolumeClaim": map[string]interface{}{
							"claimName":   "central-claim-name",
							"createClaim": true,
							// TODO(ROX-7149): more details TBD, values files are inconsistent and require more investigation and template reading
						},
					},
					"resources": map[string]interface{}{
						"limits": map[string]interface{}{
							"cpu":    "10",
							"memory": "20",
						},
						"requests": map[string]interface{}{
							"cpu":    "30",
							"memory": "40",
						},
					},
					"serviceTLS": map[string]interface{}{
						"cert": "central-tls-spec-secret-cert-content",
						"key":  "central-tls-spec-secret-key-content",
					},
				},
				"env": map[string]interface{}{
					"offlineMode": true,
					"proxyConfig": "proxy-config-secret-content",
				},
				"imagePullSecrets": map[string]interface{}{
					"useExisting": []string{
						"image-pull-secrets-secret1",
						"image-pull-secrets-secret2",
					},
				},
				"customize": map[string]interface{}{
					"annotations": map[string]string{
						"customize-annotation1": "customize-annotation1-value",
						"customize-annotation2": "customize-annotation2-value",
					},
					"labels": map[string]string{
						"customize-label1": "customize-label1-value",
						"customize-label2": "customize-label2-value",
					},
					"envVars": map[string]string{
						"customize-env-var1": "customize-env-var1-value",
						"customize-env-var2": "customize-env-var2-value",
					},
				},
				"scanner": map[string]interface{}{
					"disable":  false,
					"replicas": int32(7),
					"autoscaling": map[string]interface{}{
						"disable":     false,
						"minReplicas": int32(6),
						"maxReplicas": int32(8),
					},
					"nodeSelector": map[string]string{
						"scanner-node-selector-label1": "scanner-node-selector-value1",
						"scanner-node-selector-label2": "scanner-node-selector-value2",
					},
					"dbNodeSelector": map[string]string{
						"scanner-db-node-selector-label1": "scanner-db-node-selector-value1",
						"scanner-db-node-selector-label2": "scanner-db-node-selector-value2",
					},
					"resources": map[string]interface{}{
						"limits": map[string]interface{}{
							"cpu":    "50",
							"memory": "60",
						},
						"requests": map[string]interface{}{
							"cpu":    "70",
							"memory": "80",
						},
					},
					"dbResources": map[string]interface{}{
						"limits": map[string]interface{}{
							"cpu":    "90",
							"memory": "100",
						},
						"requests": map[string]interface{}{
							"cpu":    "110",
							"memory": "120",
						},
					},
					"serviceTLS": map[string]interface{}{
						"cert": "scanner-tls-spec-secret-cert-content",
						"key":  "scanner-tls-spec-secret-key-content",
					},
					"dbServiceTLS": map[string]interface{}{
						"cert": "scanner-db-tls-spec-secret-cert-content",
						"key":  "scanner-db-tls-spec-secret-key-content",
					},
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			wantAsValues, err := translation.ToHelmValues(tt.want)
			require.NoError(t, err, "error in test specification: cannot translate `want` specification to Helm values")

			got, err := translate(context.Background(), tt.args.clientSet, tt.args.c)
			assert.NoError(t, err)

			assert.Equal(t, wantAsValues, got)
		})
	}
}

func makeSecret(name string, stringData map[string]string) *corev1.Secret {
	data := map[string][]byte{}
	for key, val := range stringData {
		data[key] = []byte(val)
	}
	return &corev1.Secret{
		ObjectMeta: v1.ObjectMeta{Name: name, Namespace: "stackrox"},
		Data:       data,
	}
}
