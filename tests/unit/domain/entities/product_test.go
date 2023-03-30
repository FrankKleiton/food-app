package test_entities

import (
	"food-app/domain/entities"
	"food-app/tests/testing_utils"
	"testing"
)

func TestProduct(t *testing.T) {
	t.Run("check types", func(t *testing.T) {
		product := entities.Product{}
		notFoundProduct := entities.NotFoundProduct{}

		testing_utils.AssertEqual(product.IsValid(), true, t)
		testing_utils.AssertEqual(notFoundProduct.IsValid(), false, t)
	})
}
