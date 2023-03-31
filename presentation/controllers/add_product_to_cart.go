package controllers

import (
	"food-app/domain/adapters"
	"food-app/presentation/requests"
)

type AddProductToCart struct {
	Interactor adapters.IAddProductToCart
}

func (c *AddProductToCart) AddToCart(request requests.ProductsIds) (adapters.ICart, error) {
	cart, err := c.Interactor.Execute(request.Get())

	return cart, err
}
