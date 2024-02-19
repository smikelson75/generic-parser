package handlers

import "github.com/smikelson75/parser/handlers/patterns"

type IPatternsBuffer interface {
	GetTo(start, end patterns.IPattern) (*string, error)
}