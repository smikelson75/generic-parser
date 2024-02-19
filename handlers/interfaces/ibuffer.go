package interfaces

type IBuffer interface {
	Get(pattern IPattern) (*string, error)
}