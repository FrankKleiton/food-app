package http

import (
	"encoding/json"
	"net/http"

	"food-app/domain/adapters"
	"food-app/domain/interactors"
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

		result, _ := usecase.Execute(model.Products)

		cartJson, _ := json.Marshal(result)

		// if ()

		response.Header().Set("Content-Type", "application/json")
		response.Write(cartJson)
	})

	return &router
}
