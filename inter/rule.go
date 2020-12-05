package inter

import "github.com/lanvard/support"

type Rule interface {
	Verify(value support.Value) error
}

type RuleNeedToBePresent interface {
	NeedToBePresent()
}
