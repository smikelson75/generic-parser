package parser_test

import (
	"io"
	"testing"

	"github.com/smikelson75/parser/handlers"
	"github.com/smikelson75/parser/readers"
)

func TestReaderMultiple(t *testing.T) {
	buffer := readers.FromString("abc", 60)

	err := buffer.ReadTo('\n')
	if err != nil && err != io.EOF {
		t.Fatalf("unexpected error: %v", err)
	}

	token, err := buffer.Get(handlers.NewIdentifiers())
	if err != nil && err != io.EOF {
		t.Fatalf("unexpected error: %v", err)
	}

	if token == nil {
		t.Fatalf("expected token to be returned")
	}
}
