package gateways

import (
	"food-app/domain/adapters"
	"food-app/domain/entities"
)

type MemoryCartGateway struct {
}

func (g *MemoryCartGateway) GetFilledCart() adapters.ICart {
	return &entities.NotFoundCart{}
}
