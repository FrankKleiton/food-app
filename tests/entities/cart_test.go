package test_entities

import (
	"testing"

	"food-app/utils"
)

func TestProduct(t *testing.T) {
	t.Run("Add product to cart", func(t *testing.T) {
		products, cart, _, _, teardown := utils.SetupProducts(t)
		defer teardown()

		products[0].EXPECT().GetId().AnyTimes().Return("1")

		cart.AddItem(products[0])

		utils.AssertEqual(len(cart.GetItems()), 1, t)
	})

	t.Run("Add multiple unique product to cart", func(t *testing.T) {
		products, cart, _, _, teardown := utils.SetupProducts(t)
		defer teardown()

		products[0].EXPECT().GetId().AnyTimes().Return("1")

		for i := 0; i < 3; i++ {
			cart.AddItem(products[0])
		}

		found := cart.GetItem(products[0].GetId())

		utils.AssertEqual(found.GetAmount(), 3, t)
	})

	t.Run("Add multiple products to cart", func(t *testing.T) {
		products, cart, _, _, teardown := utils.SetupProducts(t)
		defer teardown()

		products[0].EXPECT().GetId().AnyTimes().Return("1")
		products[1].EXPECT().GetId().AnyTimes().Return("2")
		products[2].EXPECT().GetId().AnyTimes().Return("3")

		cart.AddItem(products[0])

		for i := 0; i < 3; i++ {
			cart.AddItem(products[i])
		}

		utils.AssertEqual(len(cart.GetItems()), 3, t)
	})

	t.Run("Calculate cart total", func(t *testing.T) {
		products, cart, _, _, teardown := utils.SetupProducts(t)
		defer teardown()

		products[0].EXPECT().GetPrice().AnyTimes().Return(100.00)
		products[1].EXPECT().GetPrice().AnyTimes().Return(200.00)
		products[2].EXPECT().GetPrice().AnyTimes().Return(300.00)

		cart.AddItem(products[0])

		for i := 0; i < 3; i++ {
			cart.AddItem(products[i])
		}

		utils.AssertEqual(cart.GetTotal(), 700.00, t)
	})
}
