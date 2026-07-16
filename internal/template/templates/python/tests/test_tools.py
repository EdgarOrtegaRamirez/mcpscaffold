"""Tests for {{PROJECT_NAME}} tools."""
import pytest


def test_import_tools():
    """Test that the tools module can be imported."""
    import {{PROJECT_NAME_LOWER}}.tools  # noqa: F401


def test_echo():
    """Test the echo tool returns input."""
    from {{PROJECT_NAME_LOWER}}.tools import echo_tool
    result = echo_tool("test message")
    assert result.text == "test message"