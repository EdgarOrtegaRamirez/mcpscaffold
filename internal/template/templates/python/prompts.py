"""Prompts for {{PROJECT_NAME}}."""

from mcp.types import TextContent


def get_greeting_prompt(name: str) -> list[TextContent]:
    """Generate a greeting prompt.

    Args:
        name: The name to greet.

    Returns:
        A list of text content items.
    """
    return [
        TextContent(type="text", text=f"Hello, {name}!")
    ]