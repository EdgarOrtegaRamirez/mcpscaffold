# {{PROJECT_NAME}}

{{DESCRIPTION}}

A production-ready MCP (Model Context Protocol) server built with {{LANGUAGE}}.

## Features

- {{TRANSPORT}} transport layer
- Type-safe with {{LANGUAGE}} type hints
- Structured logging
- Graceful shutdown handling
- Full test coverage
- CI/CD pipeline

## Quick Start

```bash
# Run with {{TRANSPORT}} transport (default)
python -m {{PROJECT_NAME_LOWER}}
```

## Installation

```bash
pip install -r requirements.txt
```

## Usage

### From a client

```python
from mcp import ClientSession
import asyncio

async def main():
    async with ClientSession() as session:
        await session.initialize()
        tools = await session.list_tools()
        print(f"Available tools: {[t.name for t in tools]}")

asyncio.run(main())
```

## Project Structure

```
{{PROJECT_NAME_LOWER}}/
├── {{PROJECT_NAME_LOWER}}/
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
└── .github/
    └── workflows/
        └── ci.yml
```

## Development

```bash
# Install dev dependencies
pip install -e ".[dev]"

# Run tests
pytest tests/ -v

# Run linter
ruff check {{PROJECT_NAME_LOWER}}/

# Format code
ruff format {{PROJECT_NAME_LOWER}}/
```

## License

MIT