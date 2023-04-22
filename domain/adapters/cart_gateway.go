package adapters

type ICartGateway interface {
	GetActiveCart(c chan ICart)
	SaveCart(cart ICart) bool
	UpdateCart(cart ICart) bool
}
