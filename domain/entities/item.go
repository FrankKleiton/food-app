package entities

import "food-app/domain/adapters"

type Item struct {
	Amount  int
	Product adapters.IProduct
}

func (p *Item) Increment(amount int) {
	p.Amount = p.Amount + amount
}

func (p Item) GetAmount() int {
	return p.Amount
}

func (p Item) GetId() string {
	prod := p.Product

	return prod.GetId()
}

func (p Item) GetTotal() float64 {
	return (p.Product).GetPrice() * float64(p.Amount)
}
