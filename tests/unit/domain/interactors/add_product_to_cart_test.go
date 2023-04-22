package interactors

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"food-app/domain/entities"
	"food-app/domain/interactors"
	"food-app/tests/mocks"
	"food-app/tests/testing_utils"
)

var productGateway *mocks.MockIProductGateway
var cartGateway *mocks.MockICartGateway
var products []*mocks.MockIProduct
var ctrl *gomock.Controller

func Before(callback func(t *testing.T)) func(t *testing.T) {
	ctrl = gomock.NewController(&testing.T{})
	productGateway = mocks.NewMockIProductGateway(ctrl)
	cartGateway = mocks.NewMockICartGateway(ctrl)
	products = testing_utils.MakeProducts(ctrl)
	defer ctrl.Finish()

	return callback
}

func TestAddProductToCart(t *testing.T) {
	t.Run("Add product to new cart", Before(func(t *testing.T) {
		productIds := []string{"0", "1"}

		cartGateway := mocks.CartGatewayMock{}

		productGateway.EXPECT().FindById(productIds[0]).Return(products[0])
		productGateway.EXPECT().FindById(productIds[1]).Return(products[1])

		sut := interactors.AddProductToCart{ProductGateway: productGateway, CartGateway: cartGateway}

		result, _ := sut.Execute(productIds)

		testing_utils.AssertEqual(len(result.GetItems()), 2, t)

		for index, product := range result.GetItems() {
			testing_utils.AssertEqual(product.GetId(), productIds[index], t)
		}
	}))

	t.Run("save cart", Before(func(t *testing.T) {
		productIds := []string{"1"}

		cartGateway := mocks.CartGatewayMock{}
		productGateway.EXPECT().FindById(gomock.Any()).Return(products[1])

		sut := interactors.AddProductToCart{ProductGateway: productGateway, CartGateway: cartGateway}

		sut.Execute(productIds)

		cart := entities.Cart{}

		cart.AddItem(products[0])

		productGateway = mocks.NewMockIProductGateway(ctrl)
		cartGateway = mocks.CartGatewayMock{}

		productGateway.EXPECT().FindById(gomock.Any()).Return(products[1])

		another := interactors.AddProductToCart{ProductGateway: productGateway, CartGateway: cartGateway}

		another.Execute(productIds)
	}))

	t.Run("Add product to existing cart", Before(func(t *testing.T) {
		id := "1"
		cartGateway := mocks.CartGatewayMock{}

		productGateway.EXPECT().FindById(id).Return(nil)
		sut := interactors.AddProductToCart{ProductGateway: productGateway, CartGateway: cartGateway}

		result, error := sut.Execute([]string{id})

		testing_utils.AssertTrue(reflect.TypeOf(result) == reflect.TypeOf(&entities.NotFoundCart{}), t)
		testing_utils.AssertEqual(error.Error(), "cannot add null value to cart", t)
	}))
}
