package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Server struct {
	name    string
	version string
	mcp     *mcp.Server
}

func NewServer(name, version string) *Server {
	srv := &Server{
		name:    name,
		version: version,
	}
	srv.setupMCP()
	return srv
}

func (s *Server) setupMCP() {
	s.mcp = mcp.NewServer(&mcp.ServerInfo{
		Name:    s.name,
		Version: s.version,
	})

	// Register echo tool
	s.mcp.RegisterTool(&mcp.Tool{
		Name:        "echo",
		Description: "Echo back the input text",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"message": map[string]interface{}{
					"type":        "string",
					"description": "Text to echo back",
				},
			},
			"required": []string{"message"},
		},
	}, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		msg, _ := req.Arguments["message"].(string)
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				{Type: "text", Text: msg},
			},
		}, nil
	})
}

func (s *Server) RunStdio(ctx context.Context) error {
	log.Println("Connecting to stdio transport...")
	transport := mcp.NewStdioTransport()
	if err := s.mcp.Connect(ctx, transport); err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	<-ctx.Done()
	return ctx.Err()
}

func (s *Server) RunSSE(ctx context.Context, port int) error {
	log.Printf("Starting SSE server on port %d...", port)
	transport := mcp.NewSSETransport(&mcp.SSEOptions{
		Port: port,
	})
	if err := s.mcp.Connect(ctx, transport); err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	<-ctx.Done()
	return ctx.Err()
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	srv := NewServer("{{PROJECT_NAME_UPPER}}", "0.1.0")

	transport := os.Getenv("MCP_TRANSPORT")
	if transport == "" {
		transport = "stdio"
	}
	port := 8080
	if p := os.Getenv("MCP_PORT"); p != "" {
		fmt.Sscanf(p, "%d", &port)
	}

	log.Printf("Starting {{PROJECT_NAME}} (%s transport)", transport)

	var err error
	switch transport {
	case "stdio":
		err = srv.RunStdio(ctx)
	case "sse":
		err = srv.RunSSE(ctx, port)
	case "http":
		log.Println("HTTP transport not yet implemented")
		err = ctx.Err()
	default:
		log.Fatalf("Unknown transport: %s", transport)
	}

	if err != nil && err != context.Canceled {
		log.Fatalf("Server error: %v", err)
	}
}