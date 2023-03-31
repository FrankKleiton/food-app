package http

import (
	"net/http"
)

type Router struct {
	RoutesList RouteCallback
}

func (r *Router) AddRoute(method string, path string, handler func(response http.ResponseWriter, request *http.Request)) {
	if r.RoutesList == nil {
		r.RoutesList = make(RouteCallback)
	}
	r.RoutesList[method+" "+path] = handler
}

func (r *Router) Route(response http.ResponseWriter, request *http.Request) {
	method := request.Method
	path := request.URL.Path
	r.RoutesList[method+" "+path](response, request)
}

type RouteCallback map[string]func(http.ResponseWriter, *http.Request)
