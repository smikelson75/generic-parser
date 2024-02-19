package handlers

import "github.com/smikelson75/parser/tokens"

type SingleLineComment struct {
	Handler
}

func NewSingleLineComment() *SingleLineComment {
	return &SingleLineComment{
		Handler: *NewHandler(`^\/\/.*`, tokens.COMMENT),
	}
}