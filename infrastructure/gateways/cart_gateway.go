package gateways

// type CartGateway struct {
// }

// func (g *CartGateway) GetActiveCart() adapters.ICart {
// 	connectionString := os.Getenv("CONNECTION_STRING")

// 	db, err := sql.Open("postgres", connectionString)

// 	if err != nil {
// 		// return entities.MakeErrorProduct("ServerError", err.Error())
// 		// TODO: return error cart
// 		return nil
// 	}

// 	query := "SELECT id, name, price, description, image FROM products WHERE id = $1 LIMIT 1"

// 	result := db.QueryRow(query, id)

// 	product := entities.Product{}

// 	err = result.Scan(&product.Id, &product.Name, &product.Price, &product.Description, &product.Image)

// 	if err != nil {
// 		result := err.Error()

// 		if result == "sql: no rows in result set" {
// 			return entities.MakeErrorProduct("NotFound")
// 		}

// 		return entities.MakeErrorProduct("ServerError", result)
// 	}

// 	return &product
// 	return nil
// }

// func (g *CartGateway) SaveCart(cart adapters.ICart) bool {
// 	return true
// }

// func (g *CartGateway) UpdateCart(cart adapters.ICart) bool {
// 	return true
// }
