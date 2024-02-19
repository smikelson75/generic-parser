package handlers

import (
	"io"

	"github.com/smikelson75/parser/handlers/interfaces"
	"github.com/smikelson75/parser/tokens"
)

type Handler struct {
	pattern string
	kind    tokens.TokenType
}

func NewHandler(pattern string, kind tokens.TokenType) *Handler {
	return &Handler{
		pattern: pattern,
		kind:    kind,
	}
}

func (h Handler) Get(buffer interfaces.IReader) (*tokens.Token, error) {
	value, err := buffer.Get(h)
	if err != nil && err != io.EOF {
		return nil, err
	}

	if value == nil {
		return nil, err
	}

	return tokens.NewToken(h.kind, *value), err
}

func (h Handler) Pattern() string {
	return h.pattern
}

func (h Handler) Kind() tokens.TokenType {
	return h.kind
}
