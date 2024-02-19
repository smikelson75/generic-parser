package patterns

type Pattern struct {
	pattern string
}

func NewPattern(pattern string) *Pattern {
	return &Pattern{
		pattern: pattern,
	}
}

func (p Pattern) Pattern() string {
	return p.pattern
}