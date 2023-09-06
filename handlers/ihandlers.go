package handlers

import "github.com/smikelson75/parser/tokens"

type IHandler interface {
	Get(buffer IBuffer) (*tokens.Token, error)
	Pattern() string
}
