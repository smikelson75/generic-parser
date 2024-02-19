package handlers

import (
	"github.com/smikelson75/parser/handlers/patterns"
	"github.com/smikelson75/parser/readers"
)

type IReader interface {
	Get(pattern readers.TokenPattern) (*string, error)
	GetTo(start, end patterns.IPattern) (*string, error)
}
