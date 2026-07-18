package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/EdgarOrtegaRamirez/mcpscaffold/internal/template"
	"github.com/spf13/cobra"
)

// NewRootCommand creates the root CLI command
func NewRootCommand(version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "mcpscaffold",
		Short: "Scaffold production-ready MCP server projects",
		Long: `mcpscaffold generates complete MCP server projects with proper structure,
dependencies, tests, and CI/CD pipelines. Supports Python, TypeScript, and Go.`,
		Version: version,
	}

	rootCmd.AddCommand(newScaffoldCmd())
	rootCmd.AddCommand(newListCmd())

	return rootCmd
}

func newScaffoldCmd() *cobra.Command {
	var (
		language    string
		transport   string
		name        string
		outputDir   string
		tools       int
		resources   int
		prompts     int
		description string
	)

	cmd := &cobra.Command{
		Use:   "scaffold [project-name]",
		Short: "Generate a new MCP server project",
		Long: `Generate a complete MCP server project with the selected language,
transport layer, and features.

Available languages: python, typescript, go
Available transports: stdio, sse, http
Features: tools, resources, prompts

Examples:
  mcpscaffold scaffold my-server --language python --transport stdio
  mcpscaffold scaffold api-server --language typescript --transport sse --tools 3
  mcpscaffold scaffold grpc-client --language go --transport http --tools 5`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name = args[0]

			// Validate language
			switch language {
			case "python", "typescript", "go":
				// valid
			default:
				return fmt.Errorf("unsupported language %q: use python, typescript, or go", language)
			}

			// Validate transport
			switch transport {
			case "stdio", "sse", "http":
				// valid
			default:
				return fmt.Errorf("unsupported transport %q: use stdio, sse, or http", transport)
			}

			// Validate feature counts
			if tools < 0 || resources < 0 || prompts < 0 {
				return fmt.Errorf("feature counts must be non-negative")
			}

			// Set output directory
			if outputDir == "" {
				outputDir = filepath.Join(".", name)
			}

			// Set description
			if description == "" {
				description = fmt.Sprintf("An MCP server for %s", name)
			}

			cfg := &template.Config{
				ProjectName: name,
				Description: description,
				Language:    language,
				Transport:   transport,
				OutputDir:   outputDir,
				Tools:       tools,
				Resources:   resources,
				Prompts:     prompts,
			}

			return template.Generate(cfg)
		},
	}

	cmd.Flags().StringVarP(&language, "language", "l", "python", "Programming language (python, typescript, go)")
	cmd.Flags().StringVarP(&transport, "transport", "t", "stdio", "Transport layer (stdio, sse, http)")
	cmd.Flags().StringVarP(&outputDir, "output", "o", "", "Output directory (default: project name)")
	cmd.Flags().IntVar(&tools, "tools", 1, "Number of example tools to generate")
	cmd.Flags().IntVar(&resources, "resources", 1, "Number of example resources to generate")
	cmd.Flags().IntVar(&prompts, "prompts", 0, "Number of example prompts to generate")
	cmd.Flags().StringVarP(&description, "description", "d", "", "Project description")
	cmd.Flags().String("author", "EdgarOrtegaRamirez", "Project author")

	_ = cmd.Flags().MarkHidden("author") // reserved for future use

	return cmd
}

func newListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list-languages",
		Short: "List supported languages and transports",
		Long:  "Display the list of supported programming languages and transport layers.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Supported languages:")
			for _, lang := range template.SupportedLanguages() {
				fmt.Printf("  - %s\n", lang)
			}

			fmt.Println("\nSupported transports:")
			for _, t := range template.SupportedTransports() {
				fmt.Printf("  - %s\n", t)
			}

			fmt.Println("\nAvailable templates:")
			for _, t := range template.SupportedTemplates() {
				fmt.Printf("  - %s\n", t)
			}
		},
	}
}

// Helper: ensure directory exists
func ensureDir(dir string) error {
	return os.MkdirAll(dir, 0o755)
}

// writeFile writes content to a file within the output directory
func writeFile(dir, name, content string) error {
	path := filepath.Join(dir, name)
	if err := ensureDir(filepath.Dir(path)); err != nil {
		return fmt.Errorf("creating directory for %s: %w", path, err)
	}
	return os.WriteFile(path, []byte(content), 0o644)
}

// generateSampleTools creates sample tool implementations
func generateSampleTools(name string, transport string, count int) []string {
	var tools []string

	toolNames := []string{"echo", "sum", "greet", "reverse", "upper", "lower", "length", "timestamp"}

	for i := 0; i < count && i < len(toolNames); i++ {
		tools = append(tools, fmt.Sprintf("_tool_%s", toolNames[i]))
	}

	return tools
}

// sanitizeProjectName converts a project name to a valid identifier
func sanitizeProjectName(name string) string {
	return strings.NewReplacer(" ", "-", "_", "-").Replace(strings.ToLower(name))
}
