package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const testBinary = "/root/workspace/mcpscaffold/mcpscaffold"

func init() {
	// Build the binary if it doesn't exist or source files are newer
	if needBuild() {
		cmd := exec.Command("go", "build", "-o", testBinary, ".")
		cmd.Dir = "/root/workspace/mcpscaffold"
		if err := cmd.Run(); err != nil {
			panic("failed to build test binary: " + err.Error())
		}
	}
}

func needBuild() bool {
	info, err := os.Stat(testBinary)
	if err != nil {
		return true
	}
	modTime := info.ModTime()

	srcFiles := []string{
		"/root/workspace/mcpscaffold/main.go",
		"/root/workspace/mcpscaffold/go.mod",
		"/root/workspace/mcpscaffold/internal/cli/root.go",
		"/root/workspace/mcpscaffold/internal/template/generate.go",
	}
	for _, f := range srcFiles {
		info, err := os.Stat(f)
		if err != nil {
			continue
		}
		if info.ModTime().After(modTime) {
			return true
		}
	}
	return false
}

func runCmd(args ...string) (*exec.Cmd, error) {
	cmd := exec.Command(testBinary, args...)
	return cmd, nil
}

func TestScaffoldPython(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "test-py")

	cmd, err := runCmd("scaffold", "test-py", "--language", "python", "--output", dir, "--description", "Test Python")
	if err != nil {
		t.Fatal(err)
	}
	if err := cmd.Run(); err != nil {
		t.Fatalf("scaffold failed: %v", err)
	}

	// Check key files exist
	files := []string{"README.md", "server.py", "requirements.txt", "pyproject.toml", "tools.py", ".gitignore", "AGENTS.md"}
	for _, f := range files {
		path := filepath.Join(dir, f)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("file %s not found", f)
		}
	}

	// Check template variables are replaced
	content, _ := os.ReadFile(filepath.Join(dir, "README.md"))
	if !strings.Contains(string(content), "Test Python") {
		t.Error("description not found in README")
	}
	if !strings.Contains(string(content), "test-py") {
		t.Error("project name not found in README")
	}
}

func TestScaffoldGo(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "test-go")

	cmd, err := runCmd("scaffold", "test-go", "--language", "go", "--output", dir, "--description", "Test Go")
	if err != nil {
		t.Fatal(err)
	}
	if err := cmd.Run(); err != nil {
		t.Fatalf("scaffold failed: %v", err)
	}

	// Check key files exist
	files := []string{"README.md", "main.go", "go.mod", "LICENSE", "AGENTS.md", ".gitignore"}
	for _, f := range files {
		path := filepath.Join(dir, f)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("file %s not found", f)
		}
	}

	// Check go.mod has correct module name
	content, _ := os.ReadFile(filepath.Join(dir, "go.mod"))
	if !strings.Contains(string(content), "test-go") {
		t.Error("module name not found in go.mod")
	}
}

func TestScaffoldTypeScript(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "test-ts")

	cmd, err := runCmd("scaffold", "test-ts", "--language", "typescript", "--output", dir, "--description", "Test TS")
	if err != nil {
		t.Fatal(err)
	}
	if err := cmd.Run(); err != nil {
		t.Fatalf("scaffold failed: %v", err)
	}

	// Check key files exist
	files := []string{"README.md", "index.ts", "package.json", "tsconfig.json", "tools.ts", "LICENSE", "AGENTS.md"}
	for _, f := range files {
		path := filepath.Join(dir, f)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("file %s not found", f)
		}
	}
}

func TestListLanguages(t *testing.T) {
	cmd, err := runCmd("list-languages")
	if err != nil {
		t.Fatal(err)
	}
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("list-languages failed: %v", err)
	}

	output := string(out)
	if !strings.Contains(output, "python") {
		t.Error("python not listed")
	}
	if !strings.Contains(output, "typescript") {
		t.Error("typescript not listed")
	}
	if !strings.Contains(output, "go") {
		t.Error("go not listed")
	}
}

func TestInvalidLanguage(t *testing.T) {
	cmd, err := runCmd("scaffold", "test", "--language", "rust")
	if err != nil {
		t.Fatal(err)
	}
	err = cmd.Run()
	if err == nil {
		t.Error("expected error for unsupported language")
	}
}

func TestInvalidTransport(t *testing.T) {
	cmd, err := runCmd("scaffold", "test", "--transport", "websocket")
	if err != nil {
		t.Fatal(err)
	}
	err = cmd.Run()
	if err == nil {
		t.Error("expected error for unsupported transport")
	}
}

func TestNegativeFeatureCount(t *testing.T) {
	cmd, err := runCmd("scaffold", "test", "--tools", "-1")
	if err != nil {
		t.Fatal(err)
	}
	err = cmd.Run()
	if err == nil {
		t.Error("expected error for negative feature count")
	}
}