package chartutil

import (
	"fmt"

	"helm.sh/helm/v3/pkg/chart"
	"sigs.k8s.io/yaml"
)

// GetDependencies reads a Chart.yaml content and returns the list of dependencies.
func GetDependencies(chartYAML []byte) ([]*chart.Dependency, error) {
	md := new(chart.Metadata)
	if err := yaml.Unmarshal(chartYAML, md); err != nil {
		return nil, fmt.Errorf("failed to parse Chart.yaml: %w", err)
	}

	return md.Dependencies, nil
}
