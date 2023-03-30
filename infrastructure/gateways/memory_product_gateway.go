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
	return &entities.NotFoundProduct{}
}

func MakeProductGateway() MemoryProductGateway {
	return MemoryProductGateway{
		Products: []entities.Product{
			{
				Id:          "1",
				Name:        "Product 1",
				Price:       1.0,
				Image:       "https://via.placeholder.com/150",
				Description: "This is a description of product 1",
			},
			{
				Id:          "2",
				Name:        "Product 2",
				Price:       2.0,
				Image:       "https://via.placeholder.com/150",
				Description: "This is a description of product 2",
			},
			{
				Id:          "3",
				Name:        "Product 3",
				Price:       3.0,
				Image:       "https://via.placeholder.com/150",
				Description: "This is a description of product 3",
			},
		},
	}
}
