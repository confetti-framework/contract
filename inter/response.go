package inter

type Response interface {
	HeaderHolder
	App() App
	SetApp(app App)
	GetContent() interface{}
	Content(content interface{})
	GetBody() string
	Body(body string) Response
	GetStatus() int
	Status(status int) Response
}
