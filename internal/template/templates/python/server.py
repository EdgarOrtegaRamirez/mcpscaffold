"""{{PROJECT_NAME}} - {{DESCRIPTION}}"""

import argparse
import logging
import os
import signal
import sys

from mcp.server import Server
from mcp.types import Tool, TextContent

# Configure logging
logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s] %(name)s: %(message)s",
)
logger = logging.getLogger(__name__)

# Initialize the MCP server
app = Server("{{PROJECT_NAME_UPPER}}")


@app.list_tools()
async def list_tools() -> list[Tool]:
    """List available tools."""
    return [
        Tool(
            name="echo",
            description="Echo back the input text",
            inputSchema={
                "type": "object",
                "properties": {
                    "message": {
                        "type": "string",
                        "description": "Text to echo back",
                    },
                },
                "required": ["message"],
            },
        ),
    ]


@app.call_tool()
async def call_tool(name: str, arguments: dict) -> list[TextContent]:
    """Handle tool calls."""
    if name == "echo":
        message = arguments.get("message", "")
        return [TextContent(type="text", text=str(message))]

    raise ValueError(f"Unknown tool: {name}")


async def run_stdio() -> None:
    """Run the server using stdio transport."""
    logger.info("Starting %s with stdio transport", "{{PROJECT_NAME}}")
    from mcp.server.stdio import stdio_server

    async with stdio_server() as (read_stream, write_stream):
        await app.run(read_stream, write_stream, app.create_initialization_options())


async def run_sse(port: int = 8080) -> None:
    """Run the server using SSE transport."""
    logger.info("Starting %s with SSE transport on port %d", "{{PROJECT_NAME}}", port)
    from mcp.server.sse import SseServerTransport
    from fastapi import FastAPI
    import uvicorn

    app_sse = FastAPI()
    sse = SseServerTransport("/messages/")

    @app_sse.get("/health")
    async def health():
        return {"status": "ok"}

    @app_sse.post("/messages/")
    async def messages(session_id: str | None = None):
        async with sse.connect_sse(
            app_sse, "/messages/", {"session_id": session_id}
        ) as (read_stream, write_stream):
            await app.run(read_stream, write_stream, app.create_initialization_options())

    config = uvicorn.Config(app_sse, host="0.0.0.0", port=port)
    server = uvicorn.Server(config)
    await server.serve()


async def run_http(port: int = 8080) -> None:
    """Run the server using HTTP transport (JSON-RPC)."""
    logger.info("Starting %s with HTTP transport on port %d", "{{PROJECT_NAME}}", port)
    from mcp.server.http import HttpServerTransport

    transport = HttpServerTransport(f":{port}/")

    async def handle_request(scope, receive, send):
        await transport.handle_request(scope, receive, send)

    import uvicorn
    from starlette.routing import Route, Router

    routes = [
        Route("/", endpoint=handle_request, methods=["POST"]),
        Route("/{path:path}", endpoint=handle_request, methods=["POST"]),
    ]
    app_http = Router(routes=routes)
    config = uvicorn.Config(app_http, host="0.0.0.0", port=port)
    server = uvicorn.Server(config)
    await server.serve()


def main() -> None:
    """Main entry point."""
    parser = argparse.ArgumentParser(description="{{DESCRIPTION}}")
    parser.add_argument(
        "--transport",
        choices=["stdio", "sse", "http"],
        default=os.environ.get("MCP_TRANSPORT", "stdio"),
        help="Transport layer to use",
    )
    parser.add_argument(
        "--port",
        type=int,
        default=int(os.environ.get("MCP_PORT", "8080")),
        help="Port for SSE/HTTP transport",
    )
    parser.add_argument(
        "--log-level",
        choices=["DEBUG", "INFO", "WARNING", "ERROR"],
        default="INFO",
        help="Logging level",
    )
    args = parser.parse_args()

    logging.getLogger().setLevel(getattr(logging, args.log_level))

    async def _main() -> None:
        loop = asyncio.get_event_loop()

        def handle_signal():
            logger.info("Received shutdown signal")
            loop.create_task(shutdown())

        loop.add_signal_handler(signal.SIGINT, handle_signal)
        loop.add_signal_handler(signal.SIGTERM, handle_signal)

        if args.transport == "stdio":
            await run_stdio()
        elif args.transport == "sse":
            await run_sse(args.port)
        elif args.transport == "http":
            await run_http(args.port)

    try:
        import asyncio
        asyncio.run(_main())
    except KeyboardInterrupt:
        logger.info("Server stopped")


if __name__ == "__main__":
    main()