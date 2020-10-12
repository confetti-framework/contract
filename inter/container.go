package inter

type Bindings map[string]interface{}
type Instances map[string]interface{}

type Container interface {
	Maker

	// Register a shared binding in the container.
	Singleton(abstract interface{}, concrete interface{})

	// Register an existing instance as shared in the container with an abstract
	Bind(abstract interface{}, concrete interface{})

	// Register an existing instance as shared in the container without an abstract
	Instance(concrete interface{}) interface{}

	// Get the container's bindings.
	Bindings() Bindings

	// Determine if the given abstract type has been bound.
	Bound(abstract string) bool

	// "Extend" an abstract type in the container.
	Extend(abstract interface{}, function func(service interface{}) interface{})
}

type Maker interface {
	// Resolve the given type from the container.
	Make(abstract interface{}) interface{}
}
