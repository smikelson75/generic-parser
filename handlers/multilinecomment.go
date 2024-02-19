package handlers

import "github.com/smikelson75/parser/tokens"

type MultiLineComment struct {
	Handler
}

func NewMultiLineComment() *MultiLineComment {
	return &MultiLineComment{
		Handler: *NewHandler(`^\/\*(.|\s)*`, tokens.COMMENT),
	}
}