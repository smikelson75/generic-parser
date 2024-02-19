package parser

import (
	"github.com/smikelson75/parser/handlers/interfaces"
	"github.com/smikelson75/parser/tokens"
)

type Parser struct {
	reader interfaces.IReader
}

func NewParser(reader interfaces.IReader) *Parser {
	return &Parser{
		reader: reader,
	}
}

func (s *Parser) Parse() ([]tokens.Token, error) {
	//repository := handlers.Create()

	return nil, nil
}
