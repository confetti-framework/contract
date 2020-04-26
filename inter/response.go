package inter

type Response interface {
	Content() string
	SetContent(content string) Response
}
