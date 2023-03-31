package main

import (
	httpServer "food-app/presentation/http"

	"food-app/domain/interactors"
	"food-app/infrastructure/gateways"
)

func main() {
	productGateway := gateways.MakeMemoryProductGateway()
	cartGateway := gateways.MemoryCartGateway{}

	addProductToCart := interactors.AddProductToCart{
		ProductGateway: productGateway,
		CartGateway:    &cartGateway,
	}

	server := httpServer.Server{
		Router: httpServer.MakeRouter(&addProductToCart),
	}

	server.Serve(5000)
}
