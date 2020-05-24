package inter

type Decorator interface {
	Decorate(route Route) Route
}

