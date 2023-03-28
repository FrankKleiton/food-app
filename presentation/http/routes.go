package http

import (
	"encoding/json"
	"food-app/domain/adapters"
	"food-app/domain/interactors"
	"net/http"
)

type AddProductToCartModel struct {
	Products []string `json:"products"`
}

func MakeRouter(
	productGateway adapters.IProductGateway,
	cartGateway adapters.ICartGateway,
) *Router {
	router := Router{}

	router.AddRoute("POST", "/cart", func(response http.ResponseWriter, request *http.Request) {
		var model AddProductToCartModel

		err := json.NewDecoder(request.Body).Decode(&model)

		if err != nil {
			println("Error: ", err.Error())
		}
		usecase := interactors.AddProductToCart{ProductGateway: productGateway, CartGateway: cartGateway}

		result := usecase.Execute(model.Products)

		json.NewEncoder(response).Encode(result)

		response.WriteHeader(http.StatusOK)
	})

	return &router
}
