package handlers

import "github.com/smikelson75/parser/readers"

type IBuffer interface {
	Get(pattern readers.TokenPattern) (*string, error)
}