package inter

type MapUrlRoutes = map[string]Route
type MapMethodRoutes map[string]MapUrlRoutes

type RouteCollection interface {
	Push(route Route) RouteCollection
	Merge(routeCollections []RouteCollection) RouteCollection
	All() []Route
	Match(request Request) Route
}

