# mcpscaffold

Generate production-ready MCP (Model Context Protocol) server projects from templates. Supports Python, TypeScript, and Go with configurable transport layers (stdio, SSE, HTTP) and feature counts.

## Features

- 🚀 Scaffold complete MCP server projects in seconds
- 🐍 Python templates with MCP SDK, Pytest, Ruff
- 🟨 TypeScript templates with Zod validation, Vitest
- 🐹 Go templates with Go SDK, testing
- ⚙️ Configurable transport layers (stdio, SSE, HTTP)
- 🔧 Configurable number of tools, resources, and prompts
- 📦 Ready-to-use CI/CD pipelines
- 📝 AGENTS.md for AI agent integration

## Install

```bash
go install github.com/EdgarOrtegaRamirez/mcpscaffold@latest
```

Or download a binary from [releases](https://github.com/EdgarOrtegaRamirez/mcpscaffold/releases).

## Quick Start

```bash
# Scaffold a Python MCP server
mcpscaffold scaffold my-server --language python --description "My MCP server"

# Scaffold a TypeScript MCP server with SSE transport
mcpscaffold scaffold api-server --language typescript --transport sse --tools 3

# Scaffold a Go MCP server with HTTP transport and custom features
mcpscaffold scaffold grpc-client --language go --transport http --tools 5 --resources 2 --prompts 1
```

## Usage

```
mcpscaffold scaffold [project-name] [flags]

Flags:
  -l, --language string   Programming language (python, typescript, go) (default "python")
  -t, --transport string  Transport layer (stdio, sse, http) (default "stdio")
  -o, --output string     Output directory (default: project name)
      --tools int         Number of example tools to generate (default 1)
      --resources int     Number of example resources to generate (default 1)
      --prompts int       Number of example prompts to generate (default 0)
  -d, --description string Project description
  -h, --help              Help for scaffold
  -v, --version           Version for mcpscaffold
```

### List supported languages and templates

```bash
mcpscaffold list-languages
```

## Generated Project Structure (Python example)

```
my-server/
├── my_server/
│   ├── __init__.py
│   ├── server.py        # Main MCP server
│   ├── tools.py         # Tool implementations
│   ├── resources.py     # Resource implementations
│   └── prompts.py       # Prompt implementations
├── tests/
│   ├── __init__.py
│   ├── test_server.py
│   └── test_tools.py
├── requirements.txt
├── pyproject.toml
├── README.md
├── LICENSE
├── AGENTS.md
└── .github/
    └── workflows/
        └── ci.yml
```

## License

MIT