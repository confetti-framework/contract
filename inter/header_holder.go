package inter

import "net/http"

type HeaderHolder interface {
	Header(key string) string
	Headers() http.Header
}
