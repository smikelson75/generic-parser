package tokens

type TokenValues interface {
	int | float64 | string | rune
}