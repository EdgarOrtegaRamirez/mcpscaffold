"""Tools for {{PROJECT_NAME}}."""

from mcp.types import TextContent


def echo_tool(message: str) -> TextContent:
    """Echo back the input text.

    Args:
        message: Text to echo back.

    Returns:
        The same message as text content.
    """
    return TextContent(type="text", text=message)