package inter

type Response interface {
	HeaderHolder
	App() App
	SetApp(app App)
	Content(content interface{})
	GetContent() interface{}
	Body() string
	SettmpBody(body string) Response
	Status(status int) Response
	GetStatus() int
}
