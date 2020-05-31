package inter

type MapMethodRoutes map[string][]Route

type RouteCollection interface {
	SetApp(app App)
	App() App
	Push(route Route) RouteCollection
	Merge(routeCollections RouteCollection) RouteCollection
	All() []Route
	Match(request Request) Route
	Prefix(prefix string) RouteCollection
	Name(name string) RouteCollection
}

