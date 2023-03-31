package adapters

type IAddProductToCart interface {
	Execute(ids []string) (ICart, error)
}
