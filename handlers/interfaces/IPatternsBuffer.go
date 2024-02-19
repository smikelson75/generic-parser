package interfaces

type IPatternsBuffer interface {
	GetTo(start, end IPattern) (*string, error)
}
