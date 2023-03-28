package adapters

type ICart interface {
	AddItem(p IProduct)
	GetItem(id string) IItem
	GetTotal() float64
	GetItems() []IItem
}
