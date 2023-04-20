package main

import (
	"database/sql"
	"os"

	"github.com/doug-martin/goqu/v9"
	"github.com/joho/godotenv"

	"food-app/domain/interactors"
	"food-app/infrastructure/gateways"
	"food-app/infrastructure/queries"
	httpServer "food-app/presentation/http"
)

func main() {
	godotenv.Load(".env")

	connectionString := os.Getenv("CONNECTION_STRING")

	db, _ := sql.Open("postgres", connectionString)
	ds := goqu.New("postgres", db)

	productGateway := gateways.MakeMemoryProductGateway()
	cartGateway := gateways.MemoryCartGateway{}

	addProductToCart := interactors.AddProductToCart{
		ProductGateway: productGateway,
		CartGateway:    &cartGateway,
	}

	productQuery := queries.ProductQuery{Database: ds}

	server := httpServer.Server{
		Router: httpServer.MakeRouter(&addProductToCart, productQuery),
	}

	server.Serve(5000)
}
