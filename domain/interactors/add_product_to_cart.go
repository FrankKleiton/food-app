package interactors

import (
	"food-app/domain/adapters"
	"food-app/domain/entities"
)

type AddProductToCart struct {
	ProductGateway adapters.IProductGateway
	CartGateway    adapters.ICartGateway
}

func (a *AddProductToCart) Execute(ids []string) adapters.ICart {
	cart, _ := a.CartGateway.GetFilledCart()

	if cart == nil {
		cart = &entities.Cart{}
	}
	for _, id := range ids {
		product, _ := a.ProductGateway.FindById(id)

		cart.AddItem(product)
	}

	return cart
}
