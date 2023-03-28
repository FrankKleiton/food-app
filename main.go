package main

import (
	"food-app/presentation/adapters"
	httpServer "food-app/presentation/http"
)

func main() {
	router := httpServer.Router{}

	router.AddRoute("GET", "/test", func(response adapters.IResponse, request adapters.IRequest) {
		// response.Write([]byte("Hello world!"))
	})

	server := httpServer.Server{
		Router: &router,
	}

	server.Serve(5000)
}
