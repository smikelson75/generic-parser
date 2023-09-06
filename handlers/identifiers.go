package handlers

import "github.com/smikelson75/parser/tokens"

type Identifiers struct {
	Handler
}

func NewIdentifiers() *Identifiers {
	return &Identifiers{
		Handler: *NewHandler("^(?:[a-zA-Z](?:[a-zA-Z0-9_]*[a-zA-Z0-9])?)", tokens.IDENTIFIER),
	}
}
