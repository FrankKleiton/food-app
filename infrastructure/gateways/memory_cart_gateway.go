package gateways

import (
	"food-app/domain/adapters"
)

type MemoryCartGateway struct {
}

func (g *MemoryCartGateway) GetFilledCart() adapters.ICart {
	return nil
}
