package queries

import (
	"food-app/domain/entities"

	"github.com/doug-martin/goqu/v9"
)

type ProductQuery struct {
	Database *goqu.Database
}

type Result map[string]any

func (q ProductQuery) GetAll() Result {
	var products []entities.Product

	err := q.Database.From("products").ScanStructs(&products)

	if err != nil {
		return Result{
			"error":      err.Error(),
			"statusCode": 404,
		}
	}

	return Result{"products": products}
}
