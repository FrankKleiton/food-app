package testing_utils

import (
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"food-app/domain/entities"
	"food-app/tests/mocks"
)

func AssertEqual(got any, want any, t *testing.T) {
	t.Helper()

	if got != want {
		t.Errorf("Expected %v, received %v", want, got)
	}
}

func SetupProducts(t *testing.T) ([]*mocks.MockIProduct, entities.Cart, *mocks.MockIProductGateway, *mocks.MockICartGateway, func()) {
	ctrl := gomock.NewController(t)
	cart := entities.Cart{}
	productGateway := mocks.NewMockIProductGateway(ctrl)
	cartGateway := mocks.NewMockICartGateway(ctrl)

	const size = 10
	products := make([]*mocks.MockIProduct, size)

	for i := 0; i < size; i++ {
		products[i] = mocks.NewMockIProduct(ctrl)
		products[i].EXPECT().GetId().AnyTimes().Return(strconv.Itoa(i))
	}

	return products, cart, productGateway, cartGateway, func() {
		ctrl.Finish()
	}
}
