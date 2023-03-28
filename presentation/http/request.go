package http

import (
	"net/http"
)

type Request struct {
	HttpRequest *http.Request
}

func (r *Request) GetMethod() string {
	return r.HttpRequest.Method
}

func (r *Request) GetUrl() string {
	return r.HttpRequest.URL.Path
}
