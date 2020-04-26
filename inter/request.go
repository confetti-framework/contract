package inter

import "net/http"

type Request interface {
	Content() string
	SetContent(content string) Request
	GetMethod() string
	SetApp(app App) Request
	GetSource() http.Request
}
