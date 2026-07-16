import { describe, it, expect } from "vitest";
import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";

describe("{{PROJECT_NAME}}", () => {
  it("should create a server instance", () => {
    const server = new McpServer({
      name: "{{PROJECT_NAME_UPPER}}",
      version: "0.1.0",
    });
    expect(server).toBeDefined();
  });
});