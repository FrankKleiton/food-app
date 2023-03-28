package presentation

import (
	"net/http"
)

type Routes struct {
	RoutesList map[string]func(http.ResponseWriter, *http.Request)
}

func (r *Routes) AddRoute(method string, path string, handler func(http.ResponseWriter, *http.Request)) {
	if r.RoutesList == nil {
		r.RoutesList = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	r.RoutesList[method+" "+path] = handler
}

func (r *Routes) Route(request *http.Request, response http.ResponseWriter) {
	method := request.Method
	path := request.URL.Path
	r.RoutesList[method+" "+path](response, request)
}
