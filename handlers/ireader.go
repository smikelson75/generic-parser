package handlers

type IReader interface {
	Next() (rune, error)
	Peek() (rune, error)
}