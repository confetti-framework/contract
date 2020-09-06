package inter

type ResponseEncoder interface {
	Can(content interface{}) bool
	Encode(content interface{}) string
}
