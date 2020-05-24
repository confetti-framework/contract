package inter

type Route interface {
	Url() string
	Method() string
	Controller() Controller
	SetPrefix(prefix string) Route
	SetDestination(destination string) Route
	SetStatus(status int) Route
	RouteOptions() RouteOptions
	Constraint() map[string]string
	SetConstraint(parameter string, regex string) Route
	SetUrl(url string) Route
}

type RouteOptions interface {
	Status() int
	Destination() string
	Prefixes() []string
}
