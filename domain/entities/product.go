package entities

type Product struct {
	Id          string
	Name        string
	Price       float64
	Image       string
	Description string
}

func (p *Product) GetId() string {
	return p.Id
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) GetImage() string {
	return p.Image
}

func (p *Product) GetDescription() string {
	return p.Description
}

func (p Product) IsValid() bool {
	return true
}
