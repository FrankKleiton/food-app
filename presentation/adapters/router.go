package adapters

import "net/http"

type IRouter interface {
	AddRoute(method string, path string, handler func(http.ResponseWriter, *http.Request))
	Route(response http.ResponseWriter, request *http.Request)
}
