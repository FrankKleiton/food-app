package presentation

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"food-app/domain/entities"
	"food-app/domain/interactors"
	"food-app/presentation/controllers"
	"food-app/presentation/requests"
	"food-app/tests/mocks"
	"food-app/tests/testing_utils"
)

var productGateway *mocks.MockIProductGateway
var cartGateway *mocks.MockICartGateway
var products []*mocks.MockIProduct

func Before(callback func(t *testing.T)) func(t *testing.T) {
	ctrl := gomock.NewController(&testing.T{})
	productGateway = mocks.NewMockIProductGateway(ctrl)
	cartGateway = mocks.NewMockICartGateway(ctrl)
	products = testing_utils.MakeProducts(ctrl)
	defer ctrl.Finish()

	return callback
}

func TestAddProductToCartController(t *testing.T) {
	t.Run("Should execute the usecase using valid product", Before(func(t *testing.T) {
		// Arrange
		// Act
		// Assert
		id := "1"
		cartGateway.EXPECT().GetActiveCart().AnyTimes().Return(nil)
		productGateway.EXPECT().FindById(id).Return(products[1])
		interactor := interactors.AddProductToCart{ProductGateway: productGateway, CartGateway: cartGateway}
		request := requests.ProductsIds{Ids: []string{id}}

		controller := controllers.AddProductToCart{Interactor: &interactor}

		result, _ := controller.AddToCart(request)

		testing_utils.AssertTrue(reflect.TypeOf(result) == reflect.TypeOf(&entities.Cart{}), t)
	}))
}
