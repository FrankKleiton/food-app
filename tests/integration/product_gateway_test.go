package integration

import (
	"database/sql"
	"os"
	"reflect"
	"testing"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/lib/pq"

	"food-app/domain/adapters"
	"food-app/domain/entities"
	"food-app/infrastructure/gateways"
	"food-app/tests/testing_utils"

	"github.com/joho/godotenv"
)

func TestProductGateway(t *testing.T) {
	godotenv.Load("../../.env")
	connectionString := os.Getenv("CONNECTION_STRING")
	db, _ := sql.Open("postgres", connectionString)
	ds := goqu.New("postgres", db)

	t.Run("Gets product by id", func(t *testing.T) {
		gateway := gateways.ProductGateway{Database: ds}

		channel := make(chan adapters.IProduct)
		go gateway.FindById("1", channel)
		product := <-channel

		testing_utils.AssertEqual(product.GetId(), "1", t)
	})

	t.Run("Gets not found product", func(t *testing.T) {
		gateway := gateways.ProductGateway{Database: ds}

		channel := make(chan adapters.IProduct)
		go gateway.FindById("2", channel)
		result := <-channel

		if reflect.TypeOf(result).Elem().Name() != "Product" {
			product := result.(*entities.ErrorProduct)

			testing_utils.AssertEqual(product.GetMessage(), "Product not found", t)
			testing_utils.AssertEqual(product.GetCode(), 404, t)
		}
	})
}
