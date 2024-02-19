package handlers

import (
	"github.com/smikelson75/parser/tokens"
)

type IBlockHandler interface {
	Get(buffer IPatternsBuffer) (*tokens.Token, error)
}