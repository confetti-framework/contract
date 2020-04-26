package inter

type MiddlewareDestination func(data Request) Response

type Pipe interface {
	Handle(data Request, next MiddlewareDestination) Response
}

