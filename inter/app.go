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
	Bind(abstract interface{}, concrete interface{})

	// Register an existing instance as shared in the container without an abstract
	Instance(concrete interface{}) interface{}

	Environment() (string, error)

	IsEnvironment(environments ...string) bool

	// The Log method gives you an instance of a logger. You can write your log messages to this instance.
	Log() Logger
}
