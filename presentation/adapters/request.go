package adapters

type IRequest interface {
	GetMethod() string
	GetUrl() string
}
