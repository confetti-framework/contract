package inter

type App interface {
	// Get the service container
	Container() *Container

	// Set the service container
	SetContainer(container Container)

	// Register a shared binding in the container.
	Singleton(abstract interface{}, concrete interface{})

	// Resolve the given type from the container.
	Make(abstract interface{}) interface{}

	// Register an existing instance as shared in the container.
	Instance(abstract interface{}, concrete interface{}) interface{}

	// Bind all of the application paths in the container.
	BindPathsInContainer(path BasePath)
}


