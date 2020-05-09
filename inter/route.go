package inter

type Route interface {
	Uri() string
	Method() string
	Controller() Controller
}
