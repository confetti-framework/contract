package inter

import (
	"net/http"
)

type Request interface {
	App() App
	SetApp(app App) Request
	Make(abstract interface{}) interface{}
	Content() string
	SetContent(content string) Request
	Method() string
	Source() http.Request
	UrlValue(key string) Value
	SetUrlValues(vars map[string]string) Request
	QueryValue(key string) Value
	QueryValues(key string) []Value
	Header(key string) string
	Headers() http.Header
	Route() Route
}
