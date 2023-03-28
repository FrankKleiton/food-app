package entities

type Product struct {
	id          string
	name        string
	price       float64
	image       string
	description string
}

func (p *Product) GetId() string {
	return p.id
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) GetPrice() float64 {
	return p.price
}

func (p *Product) GetImage() string {
	return p.image
}

func (p *Product) GetDescription() string {
	return p.description
}
