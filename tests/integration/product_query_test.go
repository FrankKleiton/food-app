package integration

import (
	"database/sql"
	"os"
	"testing"

	"food-app/infrastructure/queries"
	"food-app/tests/testing_utils"

	"github.com/doug-martin/goqu/v9"
	"github.com/joho/godotenv"
)

func TestProductQuery(t *testing.T) {
	godotenv.Load("../../.env")
	connectionString := os.Getenv("CONNECTION_STRING")

	db, _ := sql.Open("postgres", connectionString)
	ds := goqu.New("postgres", db)

	t.Run("Gets products", func(t *testing.T) {
		query := queries.ProductQuery{Database: ds}

		result := query.GetAll()

		testing_utils.AssertTrue(len(result) > 0, t)
	})
}
