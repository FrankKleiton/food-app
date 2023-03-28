package adapters

type IItem interface {
	GetAmount() int
	Increment(amount int)
	GetId() string
	GetTotal() float64
}
