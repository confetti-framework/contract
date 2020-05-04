package inter

type MapUrlRoutes = map[string]Route
type MapMethodRoutes map[string]MapUrlRoutes

type RouteCollection interface {
	SetApp(app App)
	App() App
	Push(route Route) RouteCollection
	Merge(routeCollections []RouteCollection) RouteCollection
	All() []Route
	Match(request Request) Route
}

