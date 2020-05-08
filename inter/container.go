package inter

type Bindings map[string]interface{}
type Instances map[string]interface{}

type Container interface {
	// Resolve the given type from the container.
	Make(abstract interface{}) interface{}

	// Register a shared binding in the container.
	Singleton(abstract interface{}, concrete interface{})

	// Register an existing instance as shared in the container.
	Instance(abstract interface{}, concrete interface{}) interface{}

	// // Register an existing instance as shared in the container without an abstract
	BindStruct(concrete interface{}) interface{}

	// Get the container's bindings.
	Bindings() Bindings

	Copy() Container

	// "Extend" an abstract type in the container.
	Extend(abstract interface{}, function func(service interface{}) interface{})
}
