package mocks

import (
	"food-app/domain/adapters"
)

type ProductGatewayMock struct {
	Products []adapters.IProduct
}

func (g ProductGatewayMock) FindById(id string, channel chan adapters.IProduct) {

	for _, product := range g.Products {
		if product.GetId() == id {
			channel <- product
			return
		}
	}

	channel <- nil
}
