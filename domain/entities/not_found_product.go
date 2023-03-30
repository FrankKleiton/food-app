package entities

type NotFoundProduct struct {
	Product
}

func (p *NotFoundProduct) GetId() string {
	return ""
}

func (p *NotFoundProduct) GetName() string {
	return ""
}

func (p *NotFoundProduct) GetPrice() float64 {
	return 0.0
}

func (p *NotFoundProduct) GetImage() string {
	return ""
}

func (p *NotFoundProduct) GetDescription() string {
	return ""
}

func (p NotFoundProduct) IsValid() bool {
	return false
}
