package parser

import (
	"github.com/smikelson75/parser/handlers"
	"github.com/smikelson75/parser/tokens"
)

type Parser struct {
	reader  handlers.IReader
}

func NewParser(reader handlers.IReader) *Parser {
	return &Parser{
		reader: reader,
	}
}

func (s *Parser) Parse() ([]tokens.Token, error) {
	//repository := handlers.Create()





	return nil, nil
}


