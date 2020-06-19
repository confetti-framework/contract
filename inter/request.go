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
	Source() http.Request
	Method() string
	IsMethod(method string) bool
	Path() string
	Url() string
	FullUrl() string
	All() Bag
	Value(key string) Value
	ValueOr(key string, value interface{}) Value
	Values(key string) Collection
	SetUrlValues(vars map[string]string) Request
	Header(key string) string
	Headers() http.Header
	Route() Route
}
