package inter

type MiddlewareDestination func(request Request) Response

type Pipe interface {
	Handle(request Request, next MiddlewareDestination) Response
}

