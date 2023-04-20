package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"food-app/domain/adapters"
	"food-app/infrastructure/queries"
	"food-app/presentation/controllers"
	"food-app/presentation/requests"
)

type AddProductToCartModel struct {
	Products []string `json:"products"`
}

func (m AddProductToCartModel) ContainsIds() bool {
	return len(m.Products) > 0
}

func MakeRouter(
	Interactor adapters.IAddProductToCart,
	ProductQuery queries.ProductQuery,
) *Router {
	router := Router{}

	router.AddRoute("GET", "/products", func(w http.ResponseWriter, r *http.Request) {
		result := ProductQuery.GetAll()

		JsonResponse(w, result, http.StatusOK)
	})

	router.AddRoute("POST", "/cart", func(w http.ResponseWriter, r *http.Request) {
		model := BuildModel[AddProductToCartModel](r.Body)

		if !model.ContainsIds() {
			JsonResponse(w, MakeError("We expect to receive valid product ids"), http.StatusBadRequest)
			return
		}

		request := requests.ProductsIds{Ids: model.Products}

		controller := controllers.AddProductToCart{Interactor: Interactor}

		result, _ := controller.AddToCart(request)

		JsonResponse(w, result, http.StatusOK)
	})

	return &router
}

func BuildModel[T any](source io.Reader) T {
	var destination T

	json.NewDecoder(source).Decode(&destination)

	return destination
}

func JsonResponse(w http.ResponseWriter, data any, statusCode int) {
	var responseJson []byte

	if reflect.TypeOf(data).String() == "[]byte" {
		responseJson = data.([]byte)
	} else {
		responseJson, _ = json.Marshal(data)
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	if data != nil && data != "" {
		fmt.Fprintln(w, string(responseJson))
	}
}

func MakeError(message string) map[string]string {
	return map[string]string{"error": message}
}
