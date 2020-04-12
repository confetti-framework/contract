package inter

type Rule interface {
	Passes(value string) error
}
