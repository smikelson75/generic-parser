package parser_test

import (
	"testing"

	"github.com/smikelson75/parser/tokens"
)

func TestToken(t *testing.T) {
	token := tokens.NewToken(tokens.IDENTIFIER, "go")

	if token.Value() != "go" {
		t.Fatalf("expected 'go', got '%s'", token.Value())
	}
}
