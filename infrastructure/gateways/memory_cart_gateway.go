package gateways

import (
	"food-app/domain/adapters"
)

type MemoryCartGateway struct {
}

func (g *MemoryCartGateway) GetActiveCart() adapters.ICart {
	return nil
}
