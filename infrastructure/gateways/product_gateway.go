package gateways

import (
	"food-app/domain/adapters"
	"food-app/domain/entities"
	infrastructureAdapters "food-app/infrastructure/adapters"
)

type ProductGateway struct {
	Connection infrastructureAdapters.IConnection
}

func (g ProductGateway) FindById(id string) adapters.IProduct {
	product := entities.Product{}
	var parameters = []([]interface{}){
		[]interface{}{id},
		[]interface{}{
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Description,
			&product.Image,
		},
	}

	query := "SELECT id, name, price, description, image FROM products WHERE id = $1 LIMIT 1"

	if err := g.Connection.Query(query, parameters); err != nil {

		result := err.Error()

		if result == "sql: no rows in result set" {
			return entities.MakeErrorProduct("NotFound")
		}

		return entities.MakeErrorProduct("ServerError", result)
	}

	if product.GetId() != "" {
		return &product
	}

	return entities.MakeErrorProduct("NotFound")
}
