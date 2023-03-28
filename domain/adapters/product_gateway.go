package adapters

type IProductGateway interface {
	FindById(id string) (IProduct, error)
}
