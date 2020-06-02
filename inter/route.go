package inter

type Route interface {
	Uri() string
	Method() string
	Controller() Controller
	SetPrefix(prefix string) Route
	SetDestination(destination string) Route
	SetStatus(status int) Route
	RouteOptions() RouteOptions
	Constraint() map[string]string
	SetConstraint(parameter string, regex string) Route
	SetUri(url string) Route
	Name() string
	SetName(name string) Route
	Named(pattern ...string) bool
}

type RouteOptions interface {
	Status() int
	Destination() string
	Prefixes() []string
}
