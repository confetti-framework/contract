package inter

type Response interface {
	HeaderHolder
	App() App
	Content() interface{}
	SetContent(content interface{})
	SetApp(app App)
	Body() string
	SetBody(body string) Response
	Status() int
	SetStatus(status int) Response
}
