package inter

import "net/http"

type Response interface {
	Content() string
	SetContent(content string) Response
	Status() int
	SetStatus(status int) Response
	Header(key string) string
	Headers() http.Header
}
