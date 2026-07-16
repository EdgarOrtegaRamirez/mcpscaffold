# AGENTS.md — {{PROJECT_NAME}}

## Project Overview
{{DESCRIPTION}}

## Quick Start
```bash
go build -o {{PROJECT_NAME_LOWER}} .
./{{PROJECT_NAME_LOWER}}
```

## Build & Test
```bash
# Build
go build -o {{PROJECT_NAME_LOWER}} .

# Run tests
go test ./... -v

# Lint
golangci-lint run

# Vet
go vet ./...
```

## Architecture
- `main.go` — Main entry point with MCP server
- `tests/main_test.go` — Test suite
- `.github/workflows/ci.yml` — CI/CD pipeline

## Adding a New Tool
1. Define the tool handler function
2. Register it with the MCP server
3. Add tests in `main_test.go`