# AGENTS.md — {{PROJECT_NAME}}

## Project Overview
{{DESCRIPTION}}

## Quick Start
```bash
pip install -r requirements.txt
python -m {{PROJECT_NAME_LOWER}}
```

## Build & Test
```bash
# Install dev deps
pip install -e ".[dev]"

# Run tests
pytest tests/ -v

# Lint
ruff check {{PROJECT_NAME_LOWER}}/

# Format
ruff format {{PROJECT_NAME_LOWER}}/
```

## Architecture
- `server.py` — Main MCP server with transport selection
- `tools.py` — Tool implementations
- `resources.py` — Resource handlers
- `prompts.py` — Prompt handlers
- `tests/` — Test suite
- `.github/workflows/ci.yml` — CI/CD pipeline

## Adding a New Tool
1. Add the tool in `tools.py`
2. Call `app.tool()` decorator or register via `list_tools`
3. Add tests in `tests/test_tools.py`