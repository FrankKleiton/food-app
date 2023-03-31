package adapters

type ICartGateway interface {
	GetActiveCart() ICart
	SaveCart(cart ICart) bool
	UpdateCart(cart ICart) bool
}
