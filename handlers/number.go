package handlers

import "github.com/smikelson75/parser/tokens"

type Number struct {
	Handler
}

func NewNumber() *Number {
	return &Number{
		Handler: *NewHandler(`^[0-9]+(?:\.?[0-9]*)`, tokens.NUMBER),
	}
}
