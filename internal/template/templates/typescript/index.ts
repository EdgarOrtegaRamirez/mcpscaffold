import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";
import { SSEServerTransport } from "@modelcontextprotocol/sdk/server/sse.js";
import { z } from "zod";

const server = new McpServer({
  name: "{{PROJECT_NAME_UPPER}}",
  version: "0.1.0",
});

// Register example tool
server.tool(
  "echo",
  {
    message: z.string().describe("Text to echo back"),
  },
  async ({ message }) => {
    return {
      content: [{ type: "text", text: message }],
    };
  }
);

async function runStdio(): Promise<void> {
  const transport = new StdioServerTransport();
  await server.connect(transport);
  console.error("{{PROJECT_NAME}} running on stdio transport");
}

async function runSse(port: number): Promise<void> {
  console.error(`{{PROJECT_NAME}} running on SSE transport port ${port}`);
  // SSE server would be started here with an HTTP framework
  const transport = new SSEServerTransport(`/messages/`, null as any);
  await server.connect(transport);
}

async function runHttp(port: number): Promise<void> {
  console.error(`{{PROJECT_NAME}} running on HTTP transport port ${port}`);
  // HTTP transport setup would go here
}

async function main(): Promise<void> {
  const transport = process.env.MCP_TRANSPORT || "stdio";
  const port = parseInt(process.env.MCP_PORT || "8080", 10);

  process.on("SIGINT", () => {
    console.error("Received SIGINT, shutting down...");
    process.exit(0);
  });
  process.on("SIGTERM", () => {
    console.error("Received SIGTERM, shutting down...");
    process.exit(0);
  });

  switch (transport) {
    case "sse":
      await runSse(port);
      break;
    case "http":
      await runHttp(port);
      break;
    default:
      await runStdio();
  }
}

main().catch((error) => {
  console.error("Fatal error:", error);
  process.exit(1);
});