package inter

type MiddlewareDestination = Controller

type HttpMiddleware interface {
	Handle(request Request, next MiddlewareDestination) Response
}

