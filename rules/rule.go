package rules

type Rule interface {
	Passes(value string) error
}
