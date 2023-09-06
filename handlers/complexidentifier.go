package handlers

import "github.com/smikelson75/parser/tokens"

type ComplexIdentifier struct {
	Handler
}

func NewComplexIdentifier() *ComplexIdentifier {
	return &ComplexIdentifier{
		Handler: *NewHandler(`^"(?:(?:(\\")|[^"])|"(?:("{2})|[^"])*")*"`, tokens.IDENTIFIER),
	}
}