package main

import (
	"github.com/joho/godotenv"

	"food-app/domain/interactors"
	"food-app/infrastructure/gateways"
	httpServer "food-app/presentation/http"
)

func main() {
	godotenv.Load(".env")

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
