package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"copilot-support-agent/internal/analyzer"
	"copilot-support-agent/internal/chartutil"
	"copilot-support-agent/internal/helmfetcher"

	"github.com/spf13/cobra"
)

func main() {
	var chartPath string
	var registryUsername string
	var registryPassword string

	rootCmd := &cobra.Command{
		Use:   "analyze",
		Short: "Analyze Helm chart values.yaml for breaking changes",
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := os.Stat(chartPath); os.IsNotExist(err) {
				return fmt.Errorf("chart path %s does not exist", chartPath)
			}
			gh_token := os.Getenv("GH_TOKEN")
			if gh_token == "" {
				return fmt.Errorf("GH_TOKEN environment variable is not set")
			}
			chartYamlPath := filepath.Join(chartPath, "Chart.yaml")
			chartYaml, err := os.ReadFile(chartYamlPath)
			if err != nil {
				return fmt.Errorf("failed to read Chart.yaml: %w", err)
			}

			valuesYamlPath := filepath.Join(chartPath, "values.yaml")
			valuesYaml, err := os.ReadFile(valuesYamlPath)
			if err != nil {
				return fmt.Errorf("failed to read values.yaml: %w", err)
			}

			deps, err := chartutil.GetDependencies(chartYaml)
			if err != nil {
				return fmt.Errorf("failed to get dependencies from Chart.yaml: %w", err)
			}

			if len(deps) == 0 {
				fmt.Println("No dependency defined")
				return nil
			}

			fetchOpts := helmfetcher.FetchOptions{
				Username: registryUsername,
				Password: registryPassword,
			}

			var chartInfoParts []string
			for _, dep := range deps {
				repoURL := dep.Repository
				chartName := dep.Name
				localVersion := dep.Version

				if repoURL == "" {
					fmt.Printf("Warning: dependency %s does not have a repository URL, skipping\n", chartName)
					continue
				}

				// Always resolve the latest upstream version
				fmt.Printf("Resolving latest upstream version for %s from %s...\n", chartName, repoURL)
				latestVersion, err := helmfetcher.ResolveLatestVersion(repoURL, chartName, fetchOpts)
				if err != nil {
					fmt.Printf("Warning: failed to resolve latest version for %s: %v, falling back to Chart.yaml version %s\n", chartName, err, localVersion)
					latestVersion = localVersion
				}

				fmt.Printf("Fetching upstream values for %s from %s (latest version: %s, local version: %s)...\n", chartName, repoURL, latestVersion, localVersion)
				upstreamRaw, err := helmfetcher.FetchValuesWithOptions(repoURL, chartName, latestVersion, fetchOpts)
				if err != nil {
					fmt.Printf("Warning: failed to fetch upstream values for %s: %v\n", chartName, err)
					continue
				}

				chartInfoParts = append(chartInfoParts, fmt.Sprintf(
					"### Dependency: %s\nChart: %s, Latest Upstream Version: %s, Local Version: %s, Repo: %s\n\nUpstream values.yaml (latest version %s):\n```yaml\n%s\n```",
					chartName, chartName, latestVersion, localVersion, repoURL, latestVersion, upstreamRaw,
				))
			}

			chartInfo := string(chartYaml) + "\n\n" + strings.Join(chartInfoParts, "\n\n")

			fmt.Printf("Analyzing differences with Copilot AI...\n")
			analyzeOpts := analyzer.AnalyzeOptions{
				RegistryUsername: registryUsername,
				RegistryPassword: registryPassword,
				Token:            gh_token,
			}
			report, err := analyzer.Analyze(cmd.Context(), string(valuesYaml), chartInfo, analyzeOpts)
			if err != nil {
				return fmt.Errorf("AI analysis failed: %w", err)
			}

			fmt.Println("\n--- Analysis Report ---")
			fmt.Println(report)

			return nil
		},
	}

	rootCmd.Flags().StringVar(&chartPath, "chart-path", ".", "Path to the local Helm chart directory")
	rootCmd.Flags().StringVar(&registryUsername, "registry-username", "", "Username for authenticated Helm registries (HTTP or OCI)")
	rootCmd.Flags().StringVar(&registryPassword, "registry-password", "", "Password for authenticated Helm registries (HTTP or OCI)")
	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		log.Fatal(err)
	}
}
