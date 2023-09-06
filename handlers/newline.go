package handlers

import "github.com/smikelson75/parser/tokens"

type Newline struct {
	Handler
}

func NewNewline() *Newline {
	return &Newline{
		Handler: *NewHandler(`^\r\n|\n`, tokens.NEWLINE),
	}
}