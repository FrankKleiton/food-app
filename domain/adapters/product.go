package adapters

type IProduct interface {
	GetId() string
	GetName() string
	GetPrice() float64
	GetImage() string
	GetDescription() string
}
