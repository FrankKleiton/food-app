package http

import "net/http"

func MakeRouter() *Router {
	router := Router{}

	router.AddRoute("GET", "/test", func(response http.ResponseWriter, request *http.Request) {
		// response.Write([]byte("Hello world!"))
	})

	return &router
}
