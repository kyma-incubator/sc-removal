// Code generated by reverse-kube-resource. DO NOT EDIT.

package main

import v1unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

var (
	// Unstructured "service-manager-proxy"
	serviceManagerProxyUnstructuredServiceAccount = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "service-manager-proxy",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-manager-proxy-regsecret"
	serviceManagerProxyRegsecretUnstructuredSecret = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Secret",
			"metadata": map[string]interface{}{
				"name":      "service-manager-proxy-regsecret",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-manager-proxy-config"
	serviceManagerProxyConfigUnstructuredConfigMap = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name":      "service-manager-proxy-config",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-manager-proxy"
	serviceManagerProxyUnstructuredClusterRole = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name":      "service-manager-proxy",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-manager-proxy"
	serviceManagerProxyUnstructuredClusterRoleBinding = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name":      "service-manager-proxy",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-manager-proxy-regsecretviewer"
	serviceManagerProxyRegsecretviewerUnstructuredRole = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "Role",
			"metadata": map[string]interface{}{
				"name":      "service-manager-proxy-regsecretviewer",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-manager-proxy"
	serviceManagerProxyUnstructuredRoleBinding = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "RoleBinding",
			"metadata": map[string]interface{}{
				"name":      "service-manager-proxy",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-manager-proxy"
	serviceManagerProxyUnstructuredService = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "service-manager-proxy",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-manager-proxy"
	serviceManagerProxyUnstructuredDeployment = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      "service-manager-proxy",
				"namespace": "kyma-system",
			},
		},
	}
)
