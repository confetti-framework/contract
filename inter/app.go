package inter

type App interface {
	Maker

	// Get the service container
	Container() *Container

	// Set the service container
	SetContainer(container Container)

	// Register a shared binding in the container.
	Singleton(abstract interface{}, concrete interface{})

	// Register an existing instance as shared in the container.
	Instance(abstract interface{}, concrete interface{}) interface{}

	// Register an existing instance as shared in the container without an abstract
	BindStruct(concrete interface{}) interface{}
}
