# {{PROJECT_NAME}}

{{DESCRIPTION}}

A production-ready MCP (Model Context Protocol) server built with {{LANGUAGE}}.

## Quick Start

```bash
# Install dependencies
npm install

# Run with stdio transport (default)
npm start

# Run with SSE transport
MCP_TRANSPORT=sse npm start
```

## Features

- {{TRANSPORT}} transport layer
- Type-safe with {{LANGUAGE}} type hints
- Structured logging
- Graceful shutdown handling
- Full test coverage
- CI/CD pipeline

## Usage

### From a client

```bash
# Connect via stdio
{{PROJECT_NAME_LOWER}}

# Connect via SSE
MCP_TRANSPORT=sse MCP_PORT=8080 {{PROJECT_NAME_LOWER}}
```

## Project Structure

```
{{PROJECT_NAME_LOWER}}/
├── index.ts
├── package.json
├── tsconfig.json
├── tools.ts
├── tests/
│   └── test_server.ts
├── LICENSE
└── README.md
```

## Development

```bash
# Install dependencies
npm install

# Build
npm run build

# Run tests
npm test

# Lint
npm run lint

# Format
npm run format
```

## License

MIT