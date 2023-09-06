package parser_test

import (
	"testing"

	"github.com/smikelson75/parser/readers"
)

func TestParser(t *testing.T) {
	reader := readers.FromString("go", 25)
	if reader == nil {
		t.Errorf("reader was nil")
	}
}
