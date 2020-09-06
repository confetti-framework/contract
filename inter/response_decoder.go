package inter

type ResponseEncoder interface {
	IsValid(object interface{}) bool
	Transform(object interface{}) string
}
