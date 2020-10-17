package inter

type Response interface {
	HeaderHolder
	App() App
	SetApp(app App)
	Content() string
	SetContent(content string) Response
	Status() int
	SetStatus(status int) Response
}
