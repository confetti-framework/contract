package inter

type Next = Controller

type HttpMiddleware interface {
	Handle(request Request, next Next) Response
}

