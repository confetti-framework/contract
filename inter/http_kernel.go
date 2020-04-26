package inter

import (
	"net/http"
)

type HttpKernel interface {
	Handle(request *http.Request) http.ResponseWriter
}
