package main

import (
	httpServer "food-app/presentation/http"

	"food-app/infrastructure/gateways"
)

func main() {
	productGateway := gateways.MakeProductGateway()
	cartGateway := gateways.MemoryCartGateway{}

	server := httpServer.Server{
		Router: httpServer.MakeRouter(productGateway, &cartGateway),
	}

	server.Serve(5000)
}
