package http

import (
	"encoding/json"
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

	router.AddRoute("POST", "/cart", func(response http.ResponseWriter, request *http.Request) {
		var model AddProductToCartModel

		err := json.NewDecoder(request.Body).Decode(&model)

		if err != nil {
			println("Error: ", err.Error())
		}

		if len(model.Products) == 0 {
			response.WriteHeader(http.StatusBadRequest)
			return
		}

		interactorRequest := requests.ProductsIds{Ids: model.Products}

		controller := controllers.AddProductToCart{Interactor: Interactor}

		controller.Execute(interactorRequest)
		result, _ := Interactor.Execute(model.Products)

		cartJson, _ := json.Marshal(result)

		// if ()

		response.Header().Set("Content-Type", "application/json")
		response.Write(cartJson)
	})

	return &router
}
