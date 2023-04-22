package interactors

import (
	"errors"

	"food-app/domain/adapters"
	"food-app/domain/entities"
)

type AddProductToCart struct {
	ProductGateway adapters.IProductGateway
	CartGateway    adapters.ICartGateway
}

func (a *AddProductToCart) Execute(ids []string) (adapters.ICart, error) {
	channel := make(chan adapters.ICart)
	go a.CartGateway.GetActiveCart(channel)
	result := <-channel
	cart := result

	if cart == nil {
		cart = &entities.Cart{}
	}

	for _, id := range ids {
		product := a.ProductGateway.FindById(id)

		if product == nil {
			return &entities.NotFoundCart{}, errors.New("cannot add null value to cart")
		}

		cart.AddItem(product)
	}

	if result == nil {
		a.CartGateway.SaveCart(cart)
	} else {
		a.CartGateway.UpdateCart(cart)
	}

	return cart, nil
}
