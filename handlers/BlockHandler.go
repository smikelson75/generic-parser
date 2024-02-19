package handlers

import (
	"github.com/smikelson75/parser/tokens"
	"github.com/smikelson75/parser/handlers/patterns"
	"github.com/smikelson75/parser/handlers/interfaces"
)

type BlockHandler struct {
	startPattern interfaces.IPattern
	endPattern   interfaces.IPattern
	kind         tokens.TokenType
}

func NewBlockHandler(startPattern, endPattern string, kind tokens.TokenType) *BlockHandler {
	return &BlockHandler{
		startPattern: patterns.NewPattern(startPattern),
		endPattern:   patterns.NewPattern(endPattern),
		kind:         kind,
	}
}

func (b BlockHandler) Get(buffer interfaces.IPatternsBuffer) (*tokens.Token, error) {
	value, err := buffer.GetTo(b.startPattern, b.endPattern)
	if err != nil {
		return nil, err
	}

	return tokens.NewToken(b.kind, *value), err
}

