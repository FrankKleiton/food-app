package integration

import (
	"testing"

	"food-app/infrastructure/queries"
	"food-app/tests/testing_utils"

	"github.com/joho/godotenv"
)

func TestProductQuery(t *testing.T) {
	godotenv.Load("../../.env")

	t.Run("Gets products", func(t *testing.T) {
		query := queries.ProductQuery{}

		result := query.GetAll()

		testing_utils.AssertTrue(len(result) > 0, t)
	})
}
