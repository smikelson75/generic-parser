package interfaces

type IReader interface {
	Get(pattern IPattern) (*string, error)
	GetTo(start, end IPattern) (*string, error)
}
