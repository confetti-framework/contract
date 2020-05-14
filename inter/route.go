package inter

type Route interface {
	Url() string
	Method() string
	Controller() Controller
	SetPrefix(prefix string)
	SetDestination(destination string)
	SetStatus(status int)
	RouteOptions() RouteOptions
}

type RouteOptions interface {
	Status() int
	Destination() string
	Prefixes() []string
}
