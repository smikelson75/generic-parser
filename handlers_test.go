package parser_test

import (
	"io"
	"testing"

	"github.com/smikelson75/parser/handlers"
	"github.com/smikelson75/parser/readers"
	"github.com/smikelson75/parser/tokens"
)

func TestWhitespaceHandler(t *testing.T) {
	// reader := readers.FromString(input, buffer)
	// whitespace := handlers.NewWhitespace()

	// if response, err := whitespace.Validate(" "); err != nil {
	// 	t.Fatalf("returned error: %v", err)
	// } else if !response {
	// 	t.Fatalf("expected true, got false")
	// }
}

func TestWhitespaceFailed(t *testing.T) {
	// whitespace := handlers.NewWhitespace()

	// if response, err := whitespace.Validate("a"); err != nil {
	// 	t.Fatalf("returned error: %v", err)
	// } else if response {
	// 	t.Fatalf("expected false, got true")
	// }
}

func TestIdentifier(t *testing.T) {
	// scenerios := []struct {
	// 	input    string
	// 	expected bool
	// }{
	// 	{"a", true},
	// 	{"A", true},
	// 	{"_", false},
	// 	{"1", false},
	// 	{"a1", true},
	// 	{"a1_", false},
	// 	{"a1b", true},
	// }

	// for _, scenerio := range scenerios {
	// 	identifier := handlers.NewIdentifiers()

	// 	if response, err := identifier.Validate(scenerio.input); err != nil {
	// 		t.Fatalf("returned error: %v", err)
	// 	} else if response != scenerio.expected {
	// 		t.Fatalf("input %s expected %v, got %v", scenerio.input, scenerio.expected, response)
	// 	}
	// }
}

func TestRepository(t *testing.T) {
	repository := handlers.Create()
	reader := readers.FromString(`Hello`, 80)
	err := reader.ReadTo('\n')
	if err != nil && err != io.EOF {
		t.Fatalf("returned error: %v", err)
	}

	token, err := repository.Get(reader)
	if err != nil && err != io.EOF {
		t.Fatalf("returned error: %v", err)
	}

	if token == nil {
		t.Fatalf("expected return of Token, got nil")
	}

	if token.Type() != tokens.IDENTIFIER {
		t.Fatalf("expected token type %s, got %s", tokens.TokenType(token.Type()), token.Type())
	}

	if token.Value() != "Hello" {
		t.Fatalf("expected token value %s, got %s", "Hello", token.Value())
	}

	// if response := repository.Accept(' '); response == nil {
	// 	t.Fatalf("expected not nil, got nil")
	// } else if _, ok := response.(*handlers.Whitespace); !ok {
	// 	t.Fatalf("expected handlers.Whitespace, got %T", response)
	// }

	// if response := repository.Accept('a'); response == nil {
	// 	t.Fatalf("expected not nil, got nil")
	// }	else if _, ok := response.(*handlers.Identifiers); !ok {
	// 	t.Fatalf("expected handlers.Identifiers, got %T", response)
	// }
}
