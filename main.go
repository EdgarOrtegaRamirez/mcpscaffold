package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/EdgarOrtegaRamirez/mcpscaffold/internal/cli"
)

var version = "dev"

func main() {
	rootCmd := cli.NewRootCommand(version)

	// Set version info
	rootCmd.SetVersionTemplate("mcpscaffold version {{.Version}}\n")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// ensureDir creates a directory (and parents) if it doesn't exist
func ensureDir(dir string) error {
	return os.MkdirAll(dir, 0o755)
}

// writeFile writes content to a file, creating parents as needed
func writeFile(dir, name, content string) error {
	path := filepath.Join(dir, name)
	if err := ensureDir(filepath.Dir(path)); err != nil {
		return fmt.Errorf("creating directory for %s: %w", path, err)
	}
	return os.WriteFile(path, []byte(content), 0o644)
}