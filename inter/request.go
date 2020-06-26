package inter

import (
	"github.com/lanvard/support"
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
	All() support.Map
	Value(key string) support.Value
	ValueOr(key string, value interface{}) support.Value
	Values(key string) support.Collection
	SetUrlValues(vars map[string]string) Request
	Query(key string) support.Value
	QueryOr(key string, defaultValue interface{}) support.Value
	Header(key string) string
	Headers() http.Header
	Route() Route
}
