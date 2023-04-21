package integration

import (
	"database/sql"
	"os"
	"testing"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/lib/pq"

	"food-app/infrastructure/gateways"
	"food-app/tests/testing_utils"

	"github.com/joho/godotenv"
)

func TestCartGateway(t *testing.T) {
	godotenv.Load("../../.env")
	connectionString := os.Getenv("CONNECTION_STRING")
	db, _ := sql.Open("postgres", connectionString)
	ds := goqu.New("postgres", db)

	t.Run("Get cart", func(t *testing.T) {
		gateway := gateways.CartGateway{Database: ds}

		cart := gateway.GetActiveCart()

		testing_utils.AssertEqual(cart.GetTotal(), 10.99, t)
	})
}
