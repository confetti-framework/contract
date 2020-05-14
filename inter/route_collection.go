package inter

type Routes = []Route
type MapMethodRoutes map[string]Routes

type RouteCollection interface {
	SetApp(app App)
	App() App
	Push(route Route) RouteCollection
	Merge(routeCollections RouteCollection) RouteCollection
	All() []Route
	Match(request Request) Route
	Prefix(prefix string) RouteCollection
}

