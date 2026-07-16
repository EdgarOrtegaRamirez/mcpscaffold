# mcpscaffold

AI Agent guidelines for working with the mcpscaffold project.

## Project Overview

mcpscaffold is a Go CLI tool that generates production-ready MCP (Model Context Protocol) server projects. It supports Python, TypeScript, and Go templates with configurable transport layers and features.

## Architecture

- `main.go` — Entry point, version injection
- `internal/cli/root.go` — Cobra CLI commands (scaffold, list-languages)
- `internal/template/generate.go` — Template engine using embed.FS and string replacement
- `internal/template/templates/` — Embedded template directories (python, typescript, golang)

## Building

```bash
go build -o mcpscaffold .
go test -v ./...
```

## Adding a New Language Template

1. Create a directory under `internal/template/templates/<language>/`
2. Add template files with `{{PROJECT_NAME}}`, `{{DESCRIPTION}}`, etc. placeholders
3. Add the language to `SupportedLanguages()` in `generate.go`
4. Add tests in `mcpscaffold_test.go`
5. Update README.md

## Template Variables

| Placeholder | Replacement |
|---|---|
| `{{PROJECT_NAME}}` | Project name as provided |
| `{{PROJECT_NAME_UPPER}}` | First char uppercase |
| `{{PROJECT_NAME_LOWER}}` | Project name lowercase |
| `{{DESCRIPTION}}` | Project description |
| `{{LANGUAGE}}` | Selected language |
| `{{TRANSPORT}}` | Transport layer |
| `{{MODULE_PATH}}` | Full Go module path |
| `{{MAINTAINER}}` | "EdgarOrtegaRamirez" |