"""Resources for {{PROJECT_NAME}}."""

from mcp.types import TextResourceContents


def get_status() -> TextResourceContents:
    """Get the server status."""
    return TextResourceContents(
        uri=f"status://{{PROJECT_NAME}}",
        mimeType="text/plain",
        text="Running",
    )