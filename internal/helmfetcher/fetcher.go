package helmfetcher

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/Masterminds/semver/v3"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/registry"
	"helm.sh/helm/v3/pkg/repo"
)

// FetchOptions holds optional credentials for authenticated registries.
type FetchOptions struct {
	Username string
	Password string
}

// FetchValues downloads a chart from a public HTTP(S) Helm repository,
// extracts its values.yaml, and returns the raw YAML string.
// For backward compatibility, this calls FetchValuesWithOptions with empty credentials.
func FetchValues(repoURL, chartName, version string) (string, error) {
	return FetchValuesWithOptions(repoURL, chartName, version, FetchOptions{})
}

// FetchValuesWithOptions downloads a chart from an HTTP(S) or OCI Helm repository,
// optionally using credentials, extracts its values.yaml, and returns the raw YAML string.
func FetchValuesWithOptions(repoURL, chartName, version string, opts FetchOptions) (string, error) {
	if strings.HasPrefix(repoURL, "oci://") {
		return fetchOCIValues(repoURL, chartName, version, opts)
	}
	return fetchHTTPValues(repoURL, chartName, version, opts)
}

// fetchOCIValues pulls a chart from an OCI registry and extracts values.yaml.
func fetchOCIValues(repoURL, chartName, version string, opts FetchOptions) (string, error) {
	registryClient, err := registry.NewClient()
	if err != nil {
		return "", fmt.Errorf("failed to create OCI registry client: %w", err)
	}

	// Login if credentials are provided
	if opts.Username != "" && opts.Password != "" {
		host := strings.TrimPrefix(repoURL, "oci://")
		host = strings.Split(host, "/")[0]
		err = registryClient.Login(host,
			registry.LoginOptBasicAuth(opts.Username, opts.Password),
		)
		if err != nil {
			return "", fmt.Errorf("failed to login to OCI registry %s: %w", host, err)
		}
	}

	// Build the full OCI reference: host/path/chartName:version
	ref := strings.TrimPrefix(strings.TrimRight(repoURL, "/"), "oci://") + "/" + chartName
	if version != "" {
		ref = ref + ":" + version
	}

	result, err := registryClient.Pull(ref)
	if err != nil {
		return "", fmt.Errorf("failed to pull OCI chart %s@%s: %w", chartName, version, err)
	}

	if result.Chart == nil || result.Chart.Data == nil {
		return "", fmt.Errorf("OCI pull returned no chart data for %s@%s", chartName, version)
	}

	chrt, err := loader.LoadArchive(bytes.NewReader(result.Chart.Data))
	if err != nil {
		return "", fmt.Errorf("failed to load OCI chart archive for %s@%s: %w", chartName, version, err)
	}

	return extractValues(chrt, chartName, version)
}

// fetchHTTPValues downloads a chart from an HTTP(S) Helm repository, optionally with auth.
func fetchHTTPValues(repoURL, chartName, version string, opts FetchOptions) (string, error) {
	chartURL, err := resolveChartURL(repoURL, chartName, version, opts)
	if err != nil {
		return "", fmt.Errorf("failed to resolve chart URL for %s@%s: %w", chartName, version, err)
	}

	var getterOpts []getter.Option
	if opts.Username != "" && opts.Password != "" {
		getterOpts = append(getterOpts, getter.WithBasicAuth(opts.Username, opts.Password))
	}

	httpGetter, err := getter.NewHTTPGetter(getterOpts...)
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP getter: %w", err)
	}

	data, err := httpGetter.Get(chartURL)
	if err != nil {
		return "", fmt.Errorf("failed to download chart %s@%s: %w", chartName, version, err)
	}

	chrt, err := loader.LoadArchive(data)
	if err != nil {
		return "", fmt.Errorf("failed to load chart archive for %s@%s: %w", chartName, version, err)
	}

	return extractValues(chrt, chartName, version)
}

// extractValues finds and returns the raw values.yaml content from a loaded chart.
func extractValues(chrt *chart.Chart, chartName, version string) (string, error) {
	for _, f := range chrt.Raw {
		if f.Name == "values.yaml" {
			return string(f.Data), nil
		}
	}

	// Fallback: serialize the parsed values (loses comments/ordering)
	if chrt.Values != nil {
		return fmt.Sprintf("%v", chrt.Values), nil
	}

	return "", fmt.Errorf("no values.yaml found in chart %s@%s", chartName, version)
}

// resolveChartURL fetches the repo index.yaml and returns the download URL
// for the specified chart name and version.
func resolveChartURL(repoURL, chartName, version string, opts FetchOptions) (string, error) {
	indexURL := strings.TrimRight(repoURL, "/") + "/index.yaml"

	var getterOpts []getter.Option
	if opts.Username != "" && opts.Password != "" {
		getterOpts = append(getterOpts, getter.WithBasicAuth(opts.Username, opts.Password))
	}

	httpGetter, err := getter.NewHTTPGetter(getterOpts...)
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP getter: %w", err)
	}

	indexData, err := httpGetter.Get(indexURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch index.yaml from %s: %w", indexURL, err)
	}

	tmpFile, err := os.CreateTemp("", "helm-index-*.yaml")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file for index: %w", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	if _, err := tmpFile.Write(indexData.Bytes()); err != nil {
		return "", fmt.Errorf("failed to write index to temp file: %w", err)
	}
	tmpFile.Close()

	idx, err := repo.LoadIndexFile(tmpFile.Name())
	if err != nil {
		return "", fmt.Errorf("failed to parse index.yaml: %w", err)
	}

	cv, err := idx.Get(chartName, version)
	if err != nil {
		return "", fmt.Errorf("chart %s@%s not found in index: %w", chartName, version, err)
	}

	if len(cv.URLs) == 0 {
		return "", fmt.Errorf("no download URLs for %s@%s", chartName, version)
	}

	chartURL := cv.URLs[0]
	if !strings.HasPrefix(chartURL, "http://") && !strings.HasPrefix(chartURL, "https://") {
		chartURL = strings.TrimRight(repoURL, "/") + "/" + chartURL
	}
	return chartURL, nil
}

// ResolveLatestVersion queries an HTTP(S) or OCI Helm repository and returns
// the latest semver version string for the given chart.
func ResolveLatestVersion(repoURL, chartName string, opts FetchOptions) (string, error) {
	if strings.HasPrefix(repoURL, "oci://") {
		return resolveLatestOCIVersion(repoURL, chartName, opts)
	}
	return resolveLatestHTTPVersion(repoURL, chartName, opts)
}

// resolveLatestHTTPVersion fetches index.yaml from an HTTP repo, sorts the
// chart versions by semver, and returns the highest.
func resolveLatestHTTPVersion(repoURL, chartName string, opts FetchOptions) (string, error) {
	indexURL := strings.TrimRight(repoURL, "/") + "/index.yaml"

	var getterOpts []getter.Option
	if opts.Username != "" && opts.Password != "" {
		getterOpts = append(getterOpts, getter.WithBasicAuth(opts.Username, opts.Password))
	}

	httpGetter, err := getter.NewHTTPGetter(getterOpts...)
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP getter: %w", err)
	}

	indexData, err := httpGetter.Get(indexURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch index.yaml from %s: %w", indexURL, err)
	}

	tmpFile, err := os.CreateTemp("", "helm-index-*.yaml")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file for index: %w", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	if _, err := tmpFile.Write(indexData.Bytes()); err != nil {
		return "", fmt.Errorf("failed to write index to temp file: %w", err)
	}
	tmpFile.Close()

	idx, err := repo.LoadIndexFile(tmpFile.Name())
	if err != nil {
		return "", fmt.Errorf("failed to parse index.yaml: %w", err)
	}

	entries, ok := idx.Entries[chartName]
	if !ok || len(entries) == 0 {
		return "", fmt.Errorf("chart %s not found in repository index", chartName)
	}

	return pickLatestSemver(chartName, chartVersionStrings(entries))
}

// resolveLatestOCIVersion lists tags from an OCI registry and returns the
// highest semver tag.
func resolveLatestOCIVersion(repoURL, chartName string, opts FetchOptions) (string, error) {
	registryClient, err := registry.NewClient()
	if err != nil {
		return "", fmt.Errorf("failed to create OCI registry client: %w", err)
	}

	if opts.Username != "" && opts.Password != "" {
		host := strings.TrimPrefix(repoURL, "oci://")
		host = strings.Split(host, "/")[0]
		err = registryClient.Login(host,
			registry.LoginOptBasicAuth(opts.Username, opts.Password),
		)
		if err != nil {
			return "", fmt.Errorf("failed to login to OCI registry %s: %w", host, err)
		}
	}

	ref := strings.TrimPrefix(strings.TrimRight(repoURL, "/"), "oci://") + "/" + chartName

	tags, err := registryClient.Tags(ref)
	if err != nil {
		return "", fmt.Errorf("failed to list OCI tags for %s: %w", ref, err)
	}
	if len(tags) == 0 {
		return "", fmt.Errorf("no tags found for OCI chart %s", ref)
	}

	return pickLatestSemver(chartName, tags)
}

// chartVersionStrings extracts version strings from Helm index chart versions.
func chartVersionStrings(entries repo.ChartVersions) []string {
	versions := make([]string, 0, len(entries))
	for _, cv := range entries {
		versions = append(versions, cv.Version)
	}
	return versions
}

// pickLatestSemver parses a list of version strings as semver, sorts them
// descending, and returns the highest one.
func pickLatestSemver(chartName string, raw []string) (string, error) {
	var parsed []*semver.Version
	for _, r := range raw {
		v, err := semver.NewVersion(r)
		if err != nil {
			continue // skip non-semver tags
		}
		parsed = append(parsed, v)
	}
	if len(parsed) == 0 {
		return "", fmt.Errorf("no valid semver versions found for chart %s", chartName)
	}

	sort.Sort(sort.Reverse(semver.Collection(parsed)))
	return parsed[0].Original(), nil
}

