package queries

import (
	"database/sql"
	"os"
)

type ProductQuery struct {
}

type Result map[string]any

func (q ProductQuery) GetAll() Result {
	connectionString := os.Getenv("CONNECTION_STRING")

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return Result{
			"error":      err.Error(),
			"statusCode": 500,
		}
	}

	query := "SELECT id, name, price, description, image FROM products"

	result, err := db.Query(query)

	if err != nil {
		return Result{
			"error":      err.Error(),
			"statusCode": 500,
		}
	}

	products := []Result{}

	for result.Next() {
		product := Result{
			"Id":          "",
			"Name":        "",
			"Price":       "",
			"Description": "",
			"Image":       "",
		}

		err = result.Scan(product["Id"], product["Name"], product["Price"], product["Description"], product["Image"])

		if err != nil {

			if err.Error() == "sql: no rows in result set" {
				return Result{
					"error":      "Product not found",
					"statusCode": 404,
				}
			}

			return Result{
				"error":      err.Error(),
				"statusCode": 500,
			}
		}

		products = append(products, product)
	}

	return Result{"products": products}
}
