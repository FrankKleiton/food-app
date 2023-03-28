package adapters

type ICartGateway interface {
	GetFilledCart() (ICart, error)
}
