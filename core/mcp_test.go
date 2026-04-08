package core

import (
	"testing"

	"github.com/vektah/gqlparser/v2/ast"
)

func TestTypeToolsFallbackDescription(t *testing.T) {
	m := &MCP{}
	allTools := NewLLMToolSet()

	field := &ast.FieldDefinition{
		Name: "build",
		Type: ast.NonNullNamedType("String", nil),
	}

	typeDef := &ast.Definition{
		Name:   "MyModule",
		Kind:   ast.Object,
		Fields: []*ast.FieldDefinition{field},
	}

	schema := &ast.Schema{
		Query: &ast.Definition{Name: "Query"},
		Types: map[string]*ast.Definition{
			"MyModule": typeDef,
			"Query":    {Name: "Query"},
		},
	}

	if err := m.typeTools(allTools, nil, schema, typeDef, nil); err != nil {
		t.Fatalf("typeTools returned error: %v", err)
	}

	tool, ok := allTools.Map["MyModule_build"]
	if !ok {
		t.Fatalf("expected MyModule_build tool to be registered")
	}

	if tool.Description != "MyModule build" {
		// Ensures we fall back to "TypeName fieldName" when no description is provided.
		t.Fatalf("expected fallback description, got %q", tool.Description)
	}
}
