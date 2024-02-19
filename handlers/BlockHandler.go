package handlers

import (
	"github.com/smikelson75/parser/tokens"
	"github.com/smikelson75/parser/handlers/patterns"
)

type BlockHandler struct {
	startPattern patterns.IPattern
	endPattern   patterns.IPattern
	kind         tokens.TokenType
}

func NewBlockHandler(startPattern, endPattern string, kind tokens.TokenType) *BlockHandler {
	return &BlockHandler{
		startPattern: patterns.NewPattern(startPattern),
		endPattern:   patterns.NewPattern(endPattern),
		kind:         kind,
	}
}

func (b BlockHandler) Get(buffer IPatternsBuffer) (*tokens.Token, error) {
	value, err := buffer.GetTo(b.startPattern, b.endPattern)
	if err != nil {
		return nil, err
	}

	return tokens.NewToken(b.kind, *value), err
}

