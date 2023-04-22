package adapters

type IProductGateway interface {
	FindById(id string, channel chan IProduct)
}
