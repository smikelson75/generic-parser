package handlers

import "github.com/smikelson75/parser/tokens"

type String struct {
	Handler
}

func NewString() *String {
	return &String{
		Handler: *NewHandler(`^'(?:(?:(\\')|[^'])|'(?:('{2})|[^'])*')*'`, tokens.STRING),
	}
}
