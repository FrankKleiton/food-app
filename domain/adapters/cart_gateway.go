package adapters

type ICartGateway interface {
	GetActiveCart() ICart
}
