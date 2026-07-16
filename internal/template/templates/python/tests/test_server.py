"""Tests for {{PROJECT_NAME}} server."""
import pytest


def test_import_server():
    """Test that the server module can be imported."""
    import {{PROJECT_NAME_LOWER}}.server  # noqa: F401


def test_server_has_app():
    """Test that the server app is initialized."""
    from {{PROJECT_NAME_LOWER}}.server import app
    assert app is not None


def test_echo_tool():
    """Test the echo tool."""
    from {{PROJECT_NAME_LOWER}}.tools import echo_tool
    result = echo_tool("hello")
    assert result.text == "hello"