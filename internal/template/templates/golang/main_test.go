package main

import (
	"testing"
)

func TestServerCreation(t *testing.T) {
	srv := NewServer("{{PROJECT_NAME_UPPER}}", "0.1.0")
	if srv == nil {
		t.Fatal("server should not be nil")
	}
}

func TestServerHasMCP(t *testing.T) {
	srv := NewServer("{{PROJECT_NAME_UPPER}}", "0.1.0")
	if srv.mcp == nil {
		t.Fatal("mcp server should not be nil")
	}
}
