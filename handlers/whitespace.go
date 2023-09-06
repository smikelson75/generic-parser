package handlers

import "github.com/smikelson75/parser/tokens"

type Whitespace struct {
	Handler
}

func NewWhitespace() *Whitespace {
	return &Whitespace{
		Handler: *NewHandler(`^[^\S\r\n]*`, tokens.WHITESPACE),
	}
}
