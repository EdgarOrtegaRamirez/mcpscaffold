# AGENTS.md — {{PROJECT_NAME}}

## Project Overview
{{DESCRIPTION}}

## Quick Start
```bash
npm install
npm start
```

## Build & Test
```bash
# Build
npm run build

# Run tests
npm test

# Lint
npm run lint

# Format
npm run format
```

## Architecture
- `index.ts` — Main MCP server entry point
- `tools.ts` — Tool implementations
- `tests/test_server.ts` — Test suite
- `.github/workflows/ci.yml` — CI/CD pipeline

## Adding a New Tool
1. Define the tool with zod input schema in `tools.ts`
2. Register it with the MCP server instance
3. Add tests in `tests/test_tools.ts`