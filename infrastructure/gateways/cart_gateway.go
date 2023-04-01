package gateways

import (
	"food-app/domain/adapters"
)

type CartGateway struct {
}

func (g *CartGateway) GetActiveCart() adapters.ICart {
	return nil
}

func (g *CartGateway) SaveCart(cart adapters.ICart) bool {
	return true
}

func (g *CartGateway) UpdateCart(cart adapters.ICart) bool {
	return true
}
