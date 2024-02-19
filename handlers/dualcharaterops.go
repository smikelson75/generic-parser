package handlers

import "github.com/smikelson75/parser/tokens"

type DualCharacterOps struct {
	Handler
}

func NewSpecialOperators() *DualCharacterOps {
    return &DualCharacterOps{
        Handler: *NewHandler(`(<<|>>|\+\+|!=|==)`, tokens.OPERATOR),
    }
}