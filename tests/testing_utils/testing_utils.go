package testing_utils

import (
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"

	"food-app/tests/mocks"
)

func AssertEqual(got any, want any, t *testing.T) {
	t.Helper()

	if got != want {
		t.Errorf("Expected %v, received %v", want, got)
	}
}

func AssertTrue(got any, t *testing.T) {
	t.Helper()

	AssertEqual(got, true, t)
}

func MakeProducts(ctrl *gomock.Controller) []*mocks.MockIProduct {
	const size = 10
	products := make([]*mocks.MockIProduct, size)

	for i := 0; i < size; i++ {
		products[i] = mocks.NewMockIProduct(ctrl)
		products[i].EXPECT().GetId().AnyTimes().Return(strconv.Itoa(i))
	}

	return products
}
