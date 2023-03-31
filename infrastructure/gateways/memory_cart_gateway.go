package gateways

import (
	"food-app/domain/adapters"
)

type MemoryCartGateway struct {
}

func (g *MemoryCartGateway) GetActiveCart() adapters.ICart {
	return nil
}

func (g *MemoryCartGateway) SaveCart(cart adapters.ICart) bool {
	return true
}

func (g *MemoryCartGateway) UpdateCart(cart adapters.ICart) bool {
	return true
}
