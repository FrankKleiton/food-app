package http

import (
	"encoding/json"
	"io"
	"net/http"

	"food-app/domain/adapters"
	"food-app/presentation/controllers"
	"food-app/presentation/requests"
)

type AddProductToCartModel struct {
	Products []string `json:"products"`
}

func MakeRouter(
	Interactor adapters.IAddProductToCart,
) *Router {
	router := Router{}

	router.AddRoute("POST", "/cart", func(w http.ResponseWriter, r *http.Request) {
		model, err := BuildModel[AddProductToCartModel](r.Body)

		if len(model.Products) == 0 || err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		request := requests.ProductsIds{Ids: model.Products}

		controller := controllers.AddProductToCart{Interactor: Interactor}

		controller.AddToCart(request)

		result, _ := Interactor.Execute(model.Products)

		cartJson, err := json.Marshal(result)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(cartJson)
	})

	return &router
}

func BuildModel[T any](source io.Reader) (T, error) {
	var destination T

	err := json.NewDecoder(source).Decode(&destination)

	return destination, err
}
