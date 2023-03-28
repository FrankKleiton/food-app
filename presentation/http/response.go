package http

import "net/http"

type Response struct {
	HttpResponse http.ResponseWriter
}
