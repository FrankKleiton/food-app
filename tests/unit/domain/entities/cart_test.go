package test_entities

import (
	"testing"

	"food-app/domain/entities"
	"food-app/tests/mocks"
	"food-app/tests/testing_utils"

	"github.com/golang/mock/gomock"
)

var cart entities.Cart
var products []*mocks.MockIProduct

func Before(callback func(t *testing.T)) func(t *testing.T) {
	ctrl := gomock.NewController(&testing.T{})
	products = testing_utils.MakeProducts(ctrl)
	defer ctrl.Finish()

	cart = entities.Cart{}

	return callback
}

func TestCart(t *testing.T) {
	t.Run("Add product to cart", Before(func(t *testing.T) {
		products[0].EXPECT().GetId().AnyTimes().Return("1")

		cart.AddItem(products[0])

		testing_utils.AssertEqual(len(cart.GetItems()), 1, t)
	}))

	t.Run("Add multiple unique product to cart", Before(func(t *testing.T) {

		products[0].EXPECT().GetId().AnyTimes().Return("1")

		for i := 0; i < 3; i++ {
			cart.AddItem(products[0])
		}

		found := cart.GetItem(products[0].GetId())

		testing_utils.AssertEqual(found.GetAmount(), 3, t)
	}))

	t.Run("Add multiple products to cart", Before(func(t *testing.T) {
		products[0].EXPECT().GetId().AnyTimes().Return("1")
		products[1].EXPECT().GetId().AnyTimes().Return("2")
		products[2].EXPECT().GetId().AnyTimes().Return("3")

		cart.AddItem(products[0])

		for i := 0; i < 3; i++ {
			cart.AddItem(products[i])
		}

		testing_utils.AssertEqual(len(cart.GetItems()), 3, t)
	}))

	t.Run("Calculate cart total", Before(func(t *testing.T) {
		products[0].EXPECT().GetPrice().AnyTimes().Return(100.00)
		products[1].EXPECT().GetPrice().AnyTimes().Return(200.00)
		products[2].EXPECT().GetPrice().AnyTimes().Return(300.00)

		cart.AddItem(products[0])

		for i := 0; i < 3; i++ {
			cart.AddItem(products[i])
		}

		testing_utils.AssertEqual(cart.GetTotal(), 700.00, t)
	}))
}
