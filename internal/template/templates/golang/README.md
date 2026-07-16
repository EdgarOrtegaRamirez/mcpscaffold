# {{PROJECT_NAME}}

{{DESCRIPTION}}

A production-ready MCP (Model Context Protocol) server built with Go.

## Quick Start

```bash
# Build
go build -o {{PROJECT_NAME_LOWER}} .

# Run with stdio transport
./{{PROJECT_NAME_LOWER}}
```

## Usage

### From a client

```go
import (
    "context"
    "github.com/modelcontextprotocol/go-sdk/mcp"
)

client, err := mcp.NewClient(context.Background(), mcp.StdioTransport)
if err != nil {
    log.Fatal(err)
}

if err := client.Initialize(ctx, mcp.InitializeParams{
    ProtocolVersion: mcp.Version20241007,
    ClientInfo: mcp.Implementation{
        Name:    "{{PROJECT_NAME_LOWER}}",
        Version: "0.1.0",
    },
}); err != nil {
    log.Fatal(err)
}
```

## Project Structure

```
{{PROJECT_NAME_LOWER}}/
├── main.go
├── go.mod
├── go.sum
├── README.md
├── LICENSE
└── .gitignore
```

## Development

```bash
# Run tests
go test ./... -v

# Lint
golangci-lint run

# Vet
go vet ./...
```

## License

MIT