package integration

import (
	"reflect"
	"testing"

	"food-app/domain/entities"
	"food-app/infrastructure/database"
	"food-app/infrastructure/gateways"
	"food-app/tests/testing_utils"
)

func TestProductGateway(t *testing.T) {
	t.Run("Gets product by id", func(t *testing.T) {
		connection := database.PostgresConnection{}
		gateway := gateways.ProductGateway{Connection: connection}

		product := gateway.FindById("1")

		testing_utils.AssertEqual(product.GetId(), "1", t)
	})

	t.Run("Gets not found product", func(t *testing.T) {
		connection := database.PostgresConnection{}
		gateway := gateways.ProductGateway{Connection: connection}

		result := gateway.FindById("2")

		if reflect.TypeOf(result).Elem().Name() != "Product" {
			product := result.(*entities.ErrorProduct)

			testing_utils.AssertEqual(product.GetMessage(), "Product not found", t)
			testing_utils.AssertEqual(product.GetCode(), 404, t)
		}
	})
}
