package requests

type ProductsIds struct {
	Ids []string
}

func (p *ProductsIds) Get() []string {
	return p.Ids
}
