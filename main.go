package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"copilot-support-agent/internal/analyzer"

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

			// deps, err := chartutil.GetDependencies(chartYaml)
			// if err != nil {
			// 	return fmt.Errorf("failed to get dependencies from Chart.yaml: %w", err)
			// }

			// if len(deps) == 0 {
			// 	fmt.Println("No dependency defined")
			// 	return nil
			// }

			// // For this simple CLI, analyze the first dependency found.
			// dep := deps[0]
			// repoURL := dep.Repository
			// chartName := dep.Name
			// version := dep.Version

			// if repoURL == "" {
			// 	return fmt.Errorf("dependency %s does not have a repository URL in Chart.yaml", chartName)
			// }

			// fmt.Printf("Fetching upstream values for %s from %s (version: %s)...\n", chartName, repoURL, version)
			// upstreamRaw, err := helmfetcher.FetchValues(repoURL, chartName, version)
			// if err != nil {
			// 	return fmt.Errorf("failed to fetch upstream values: %w", err)
			// }

			// chartInfo := fmt.Sprintf("Chart: %s, Version: %s, Repo: %s", chartName, version, repoURL)

			fmt.Printf("Analyzing differences with Copilot AI...\n")
			analyzeOpts := analyzer.AnalyzeOptions{
				RegistryUsername: registryUsername,
				RegistryPassword: registryPassword,
			}
			report, err := analyzer.Analyze(cmd.Context(), string(valuesYaml), string(chartYaml), analyzeOpts)
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
