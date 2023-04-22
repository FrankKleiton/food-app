package main

func main() {
	// godotenv.Load(".env")

	// connectionString := os.Getenv("CONNECTION_STRING")

	// db, _ := sql.Open("postgres", connectionString)
	// ds := goqu.New("postgres", db)

	// productGateway := gateways.MakeMemoryProductGateway()
	// cartGateway := gateways.MemoryCartGateway{}

	// addProductToCart := interactors.AddProductToCart{
	// 	ProductGateway: productGateway,
	// 	CartGateway:    &cartGateway,
	// }

	// productQuery := queries.ProductQuery{Database: ds}

	// server := httpServer.Server{
	// 	Router: httpServer.MakeRouter(&addProductToCart, productQuery),
	// }

	// server.Serve(5000)
}
