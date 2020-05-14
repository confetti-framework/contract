package inter

import (
	"net/http"
)

type Request interface {
	App() App
	SetApp(app App) Request
	Content() string
	SetContent(content string) Request
	Method() string
	Source() http.Request
	UrlValues() UrlValues
	SetUrlValues(vars map[string]string) Request
	QueryValues() UrlValues
	Header(key string) string
	Headers() http.Header
}
