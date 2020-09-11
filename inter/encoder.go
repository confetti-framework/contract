package inter

const Encoders = "Encoders"

type Encoder interface {
	IsAble(object interface{}) bool
	EncodeThrough(object interface{}, encoders []Encoder) (string, error)
}
