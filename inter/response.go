package inter

const ResponseBodyEncoder = "response_body_encoder"

type Response interface {
	HeaderHolder
	App() App
	SetApp(app App)
	Content() string
	ContentE() (string, error)
	SetContent(content string) Response
	Status() int
	SetStatus(status int) Response
}
