package chartutil

import (
	"testing"
)

func TestGetDependencies(t *testing.T) {
	chartYAML := []byte(`
apiVersion: v2
name: my-chart
version: 1.0.0
dependencies:
  - name: velero
    version: 12.0.0
    repository: https://vmware-tanzu.github.io/helm-charts
  - name: redis
    version: 17.0.0
    repository: https://charts.bitnami.com/bitnami
  - name: envoy
    version: v1.7.1
    repository: oci://docker.io/envoyproxy/gateway-helm
`)

	deps, err := GetDependencies(chartYAML)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(deps) != 3 {
		t.Fatalf("expected 3 dependencies, got %d", len(deps))
	}

	if deps[0].Name != "velero" || deps[0].Version != "12.0.0" || deps[0].Repository != "https://vmware-tanzu.github.io/helm-charts" {
		t.Errorf("unexpected dependency 0: %+v", deps[0])
	}
	if deps[1].Name != "redis" || deps[1].Version != "17.0.0" || deps[1].Repository != "https://charts.bitnami.com/bitnami" {
		t.Errorf("unexpected dependency 1: %+v", deps[1])
	}
	if deps[2].Name != "envoy" || deps[2].Version != "v1.7.1" || deps[2].Repository != "oci://docker.io/envoyproxy/gateway-helm" {
		t.Errorf("unexpected dependency 2: %+v", deps[2])
	}
}
