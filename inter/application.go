package inter

type Application interface {
	Container() *Container
	Make(abstract interface{}) interface{}
}
