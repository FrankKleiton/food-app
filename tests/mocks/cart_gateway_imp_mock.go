package mocks

import "food-app/domain/adapters"

type CartGatewayMock struct {
}

func (g CartGatewayMock) GetActiveCart(c chan adapters.ICart) {
	c <- nil
}

func (g CartGatewayMock) SaveCart(cart adapters.ICart) bool {
	return true
}
func (g CartGatewayMock) UpdateCart(cart adapters.ICart) bool {
	return true
}
