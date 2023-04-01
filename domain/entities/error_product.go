package entities

type ErrorProduct struct {
	Product

	Message string
	Code    int
}

func (p *ErrorProduct) GetId() string {
	return ""
}

func (p *ErrorProduct) GetName() string {
	return ""
}

func (p *ErrorProduct) GetPrice() float64 {
	return 0.0
}

func (p *ErrorProduct) GetImage() string {
	return ""
}

func (p *ErrorProduct) GetDescription() string {
	return ""
}

func (p ErrorProduct) IsValid() bool {
	return false
}

func (p *ErrorProduct) GetMessage() string {
	return p.Message
}

func (p *ErrorProduct) GetCode() int {
	return p.Code
}

func MakeErrorProduct(values ...string) *ErrorProduct {
	if values[0] == "NotFound" {
		return &ErrorProduct{Message: "Product not found", Code: 404}
	}

	if len(values) > 1 {
		return &ErrorProduct{Message: values[1], Code: 500}
	} else {
		return &ErrorProduct{Message: "Internal server error", Code: 500}
	}
}
