package inter

type Route interface {
	Uri() string
	Method() string
	Controller() ControllerMethod
}
