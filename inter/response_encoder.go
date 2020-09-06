package inter

type ResponseEncoder interface {
	Transformable(object interface{}) bool
	Transform(object interface{}) string
}
