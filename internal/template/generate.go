package template

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//go:embed all:templates
var templatesFS embed.FS

// SupportedLanguages returns the list of supported programming languages.
func SupportedLanguages() []string {
	return []string{"python", "typescript", "go"}
}

// SupportedTransports returns the list of supported transport layers.
func SupportedTransports() []string {
	return []string{"stdio", "sse", "http"}
}

// SupportedTemplates returns the list of available template names.
func SupportedTemplates() []string {
	return []string{
		"python",
		"typescript",
		"go",
	}
}

// Config holds the configuration for generating a new project.
type Config struct {
	ProjectName string
	Description string
	Language    string
	Transport   string
	OutputDir   string
	Tools       int
	Resources   int
	Prompts     int
}

// Generate creates a new MCP server project based on the configuration.
func Generate(cfg *Config) error {
	langDir := cfg.Language
	if langDir == "go" {
		langDir = "golang"
	}

	// Use os.DirFS on disk - more reliable than embed.FS sub-reading
	sourceDir := filepath.Join("internal", "template", "templates", langDir)
	return filepath.WalkDir(sourceDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Get relative path from source directory
		rel, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}

		// Handle .tmpl files
		targetName := rel
		if strings.HasSuffix(rel, ".tmpl") {
			targetName = strings.TrimSuffix(rel, ".tmpl")
		}

		// Read file
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading template %s: %w", rel, err)
		}

		contentStr := processTemplateVariables(string(content), cfg)
		outPath := filepath.Join(cfg.OutputDir, targetName)

		if err := ensureDir(filepath.Dir(outPath)); err != nil {
			return fmt.Errorf("creating directory for %s: %w", outPath, err)
		}
		return os.WriteFile(outPath, []byte(contentStr), 0o644)
	})
}

// ensureDir creates a directory (and parents) if it doesn't exist.
func ensureDir(dir string) error {
	return os.MkdirAll(dir, 0o755)
}

// processTemplateVariables replaces {{variable}} placeholders in template content.
func processTemplateVariables(content string, cfg *Config) string {
	result := content

	replacements := map[string]string{
		"{{PROJECT_NAME}}":       cfg.ProjectName,
		"{{PROJECT_NAME_UPPER}}": strings.ToUpper(cfg.ProjectName[:1]) + cfg.ProjectName[1:],
		"{{PROJECT_NAME_LOWER}}": strings.ToLower(cfg.ProjectName),
		"{{DESCRIPTION}}":        cfg.Description,
		"{{LANGUAGE}}":           cfg.Language,
		"{{TRANSPORT}}":          cfg.Transport,
		"{{MODULE_PATH}}":        fmt.Sprintf("github.com/EdgarOrtegaRamirez/%s", cfg.ProjectName),
		"{{MAINTAINER}}":         "EdgarOrtegaRamirez",
	}

	for placeholder, value := range replacements {
		result = strings.ReplaceAll(result, placeholder, value)
	}

	return result
}
