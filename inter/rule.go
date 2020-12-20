package inter

import "github.com/lanvard/support"

type Rule interface {
	Verify(value support.Value) error
}

type RuleWithRequirements interface {
	Rule
	Requirements() []Rule
}
