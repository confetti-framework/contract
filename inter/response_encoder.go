package inter

type ResponseEncoder interface {
	Transformable(object interface{}) bool
	TransformThrough(object interface{}, encoders []ResponseEncoder) (string, error)
}
