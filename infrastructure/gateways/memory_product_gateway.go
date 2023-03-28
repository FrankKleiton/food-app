package gateways

import (
	"food-app/domain/adapters"
	"food-app/domain/entities"
)

type MemoryProductGateway struct {
	Products []entities.Product
}

func (g MemoryProductGateway) FindById(id string) adapters.IProduct {
	for _, product := range g.Products {
		if product.GetId() == id {
			return &product
		}
	}
	return nil
}

var defaultInstance = MemoryProductGateway{
	Products: []entities.Product{
		{
			Id:    "1",
			Name:  "Product 1",
			Price: 1.0,
		},
		{
			Id:    "2",
			Name:  "Product 2",
			Price: 2.0,
		},
		{
			Id:    "3",
			Name:  "Product 3",
			Price: 3.0,
		},
	},
}
