package http

import "food-app/presentation/adapters"

func MakeRouter() *Router {
	router := Router{}

	router.AddRoute("GET", "/test", func(response adapters.IResponse, request adapters.IRequest) {
		// response.Write([]byte("Hello world!"))
	})

	return &router
}
