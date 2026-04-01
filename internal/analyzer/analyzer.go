package analyzer

import (
	"context"
	"fmt"
	"os"
	"strings"

	copilot "github.com/github/copilot-sdk/go"

	"copilot-support-agent/internal/helmfetcher"
)

const systemPrompt = `You are a Helm chart values.yaml structural diff analyzer.

Your job is to compare an UPSTREAM values.yaml (from a public Helm chart) with a LOCAL values.yaml (the user's overrides/customizations) and produce a clear report.
You can use all available tools to fetch values.yaml from a public or private Helm chart repository (HTTP/HTTPS or OCI) for a specific chart and version. For OCI registries (oci://), provide the OCI repository URL. If the registry requires authentication, pass username and password.

## Analysis Rules

1. **Structural Diff**: Identify keys present in upstream but missing from local, and vice versa.
2. **Type Changes**: Flag where the same key has different types (e.g., upstream has a map but local has a scalar, or upstream has a list but local has a string).
3. **Breaking Changes**: Mark as BREAKING if:
   - A key the local file overrides has changed type in upstream (scalar→map, map→scalar, etc.)
   - A key the local file overrides has been removed in upstream
   - A block the local file overrides has been restructured (keys moved/renamed)
4. **Missing Overrides**: Flag new upstream keys that exist within blocks the local file already overrides — the user may need to set these.
5. **Safe Changes**: Value-only changes (same key, same type, different default) are NOT breaking.

## Output Format

Produce a structured report with:
- **Summary**: One-line verdict (breaking / no breaking changes)
- **Breaking Changes**: Table with key path, change type, details
- **Missing Overrides**: New upstream keys the user might want to set
- **Info**: Other structural differences (non-breaking)

Be concise. Use markdown tables. Do not repeat the full YAML back — just reference key paths.

also use tool "create_markdown_file" to create a markdown file Readme.md to document all the options from values.yaml user can overwrite use upstream values.yaml for document.`

// FetchHelmValuesParams defines the parameters for the fetch_helm_values tool.
type FetchHelmValuesParams struct {
	RepoURL   string `json:"repo_url" jsonschema:"Helm repository URL (e.g. https://charts.bitnami.com/bitnami or oci://registry.example.com/charts)"`
	ChartName string `json:"chart_name" jsonschema:"Chart name (e.g. nginx)"`
	Version   string `json:"version" jsonschema:"Chart version (e.g. 18.1.0)"`
	Username  string `json:"username,omitempty" jsonschema:"Optional username for authenticated registries"`
	Password  string `json:"password,omitempty" jsonschema:"Optional password for authenticated registries"`
}

// ReadMeParams defines the parameters for the create_markdown_file tool.
type ReadMeParams struct {
	Content string `json:"content" jsonschema:"The raw markdown content to write to the Readme.md file"`
}

// AnalyzeOptions holds optional credentials passed from the CLI.
type AnalyzeOptions struct {
	RegistryUsername string
	RegistryPassword string
}

// Analyze sends the upstream and local values.yaml to Copilot AI for
// structural diff analysis and breaking change detection.
func Analyze(ctx context.Context, localYAML, chartInfo string, analyzeOpts AnalyzeOptions) (string, error) {
	opts := &copilot.ClientOptions{
		LogLevel: "error",
	}
	client := copilot.NewClient(opts)

	if err := client.Start(ctx); err != nil {
		return "", fmt.Errorf("failed to start Copilot client: %w", err)
	}
	defer client.Stop()

	// Define a tool generate Readme.md
	generateReadmeTool := copilot.DefineTool(
		"create_markdown_file",
		"Create a markdown file",
		func(params ReadMeParams, inv copilot.ToolInvocation) (any, error) {
			err := os.WriteFile("Readme.md", []byte(params.Content), 0644)
			return nil, err
		},
	)
	generateReadmeTool.SkipPermission = true

	// Define a tool so the AI can fetch additional chart versions if needed
	fetchTool := copilot.DefineTool(
		"fetch_helm_values",
		"Fetch values.yaml from an HTTP(S) or OCI Helm chart repository for a specific chart and version. Supports authenticated registries via optional username/password.",
		func(params FetchHelmValuesParams, inv copilot.ToolInvocation) (any, error) {
			// Use tool-provided creds first, fall back to CLI-provided creds
			username := params.Username
			password := params.Password
			if username == "" && password == "" {
				username = analyzeOpts.RegistryUsername
				password = analyzeOpts.RegistryPassword
			}
			opts := helmfetcher.FetchOptions{
				Username: username,
				Password: password,
			}
			yaml, err := helmfetcher.FetchValuesWithOptions(params.RepoURL, params.ChartName, params.Version, opts)
			if err != nil {
				return nil, fmt.Errorf("failed to fetch values: %w", err)
			}
			return yaml, nil
		},
	)
	fetchTool.SkipPermission = true

	session, err := client.CreateSession(ctx, &copilot.SessionConfig{
		Model: "gpt-4o",
		SystemMessage: &copilot.SystemMessageConfig{
			Content: systemPrompt,
		},
		OnPermissionRequest: copilot.PermissionHandler.ApproveAll,
		Tools:               []copilot.Tool{fetchTool, generateReadmeTool},
	})
	if err != nil {
		return "", fmt.Errorf("failed to create Copilot session: %w", err)
	}
	defer session.Disconnect()

	// Collect the AI response
	var response strings.Builder
	done := make(chan struct{})
	var sessionErr error

	session.On(func(event copilot.SessionEvent) {
		switch event.Type {
		case "assistant.message":
			if event.Data.Content != nil {
				response.WriteString(*event.Data.Content)
			}
		case "error":
			if event.Data.Content != nil {
				sessionErr = fmt.Errorf("copilot error: %s", *event.Data.Content)
			}
		case "session.idle":
			close(done)
		}
	})

	// Build the prompt with both YAML files
	prompt := fmt.Sprintf(`Analyze the following Helm chart values for structural differences and breaking changes.

## Chart Info
%s

## Upstream values.yaml (from public Helm repo)
Use the fetch_helm_values tool to retrieve the upstream values.yaml. The repository URL may be HTTP(S) or OCI (oci://).

## Local values.yaml (user overrides)
`+"```yaml\n%s\n```"+`

Please provide your structural diff analysis and breaking change report.`,
		chartInfo,
		localYAML,
	)

	_, err = session.Send(ctx, copilot.MessageOptions{
		Prompt: prompt,
	})
	if err != nil {
		return "", fmt.Errorf("failed to send prompt to Copilot: %w", err)
	}

	// Wait for completion
	<-done

	if sessionErr != nil {
		return "", sessionErr
	}

	return response.String(), nil
}
