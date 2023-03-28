package interactors

import (
	"testing"

	"food-app/domain/adapters"
	"food-app/domain/entities"
	"food-app/domain/interactors"
	"food-app/utils"
)

func TestAddProductToCart(t *testing.T) {
	t.Run("Add product to new cart", func(t *testing.T) {
		products, _, productGateway, cartGateway, teardown := utils.SetupProducts(t)
		defer teardown()

		productIds := []string{"0", "1"}

		cartGateway.EXPECT().GetFilledCart().AnyTimes().Return(nil)

		productGateway.EXPECT().FindById(productIds[0]).Return(products[0])
		productGateway.EXPECT().FindById(productIds[1]).Return(products[1])

		sut := interactors.AddProductToCart{ProductGateway: productGateway, CartGateway: cartGateway}

		result := sut.Execute(productIds)

		utils.AssertEqual(len(result.GetItems()), 2, t)

		for index, product := range result.GetItems() {
			utils.AssertEqual(product.GetId(), productIds[index], t)
		}
	})

	t.Run("Add product to existing cart", func(t *testing.T) {
		products, _, productGateway, cartGateway, teardown := utils.SetupProducts(t)
		defer teardown()

		productIds := []string{"0", "1"}

		var items []adapters.IItem

		items = append(items, &(entities.Item{
			Product: products[0],
			Amount:  1,
		}))

		cart := entities.Cart{
			Items: items,
		}

		cartGateway.EXPECT().GetFilledCart().AnyTimes().Return(&cart)

		productGateway.EXPECT().FindById(productIds[1]).Return(products[1])

		sut := interactors.AddProductToCart{ProductGateway: productGateway, CartGateway: cartGateway}

		result := sut.Execute([]string{"1"})

		utils.AssertEqual(len(result.GetItems()), 2, t)

		for index, product := range result.GetItems() {
			utils.AssertEqual(product.GetId(), productIds[index], t)
		}

		utils.AssertEqual(result, &cart, t)
	})
}
