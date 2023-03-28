package main

import (
	httpServer "food-app/presentation/http"
)

func main() {
	server := httpServer.Server{
		Router: httpServer.MakeRouter(),
	}

	server.Serve(5000)
}
