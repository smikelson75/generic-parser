package handlers

import (
	"fmt"
	"io"

	"github.com/smikelson75/parser/handlers/interfaces"
	"github.com/smikelson75/parser/tokens"
)

type Repository struct {
	handlers []interfaces.IHandler
}

func Create() *Repository {
	return &Repository{
		handlers: []interfaces.IHandler{
			NewMultiLineComment(),
			NewSingleLineComment(),
			NewComplexIdentifier(),
			NewString(),
			NewNumber(),
			NewIdentifiers(),
			NewWhitespace(),
			NewNewline(),
		},
	}
}

func (r *Repository) Get(reader interfaces.IReader) (*tokens.Token, error) {
	for _, source := range r.handlers {
		token, err := source.Get(reader)
		if err != nil && err != io.EOF {
			return nil, err
		}

		if token != nil {
			return token, err
		}
	}

	return nil, fmt.Errorf("no handler was found")
}
