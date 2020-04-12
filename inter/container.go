package inter

type Container interface {
	// Resolve the given type from the container.
	Make(abstract interface{}) interface{}

	// Register a shared binding in the container.
	Singleton(abstract interface{}, concrete interface{})

	// Register an existing instance as shared in the container.
	Instance(abstract interface{}, concrete interface{}) interface{}
}
