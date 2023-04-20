package gateways

import (
	"github.com/doug-martin/goqu/v9"
	_ "github.com/lib/pq"

	"food-app/domain/adapters"
	"food-app/domain/entities"
)

type ProductGateway struct {
	Database *goqu.Database
}

func (g ProductGateway) FindById(id string) adapters.IProduct {
	product := entities.Product{}

	result, _ := g.Database.From("products").Where(goqu.Ex{"id": id}).Limit(1).ScanStruct(&product)

	if !result {
		return entities.MakeErrorProduct("NotFound")
	}

	return &product
}
