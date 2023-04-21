package gateways

import (
	"food-app/domain/adapters"
	"food-app/domain/entities"

	"github.com/doug-martin/goqu/v9"
)

type CartGateway struct {
	Database *goqu.Database
}

type ProductModel struct {
	Id          string
	Name        string
	Price       float64
	Image       string
	Description string
	Amount      int
}

type CartModel struct {
	Id string
}

func (g *CartGateway) GetActiveCart() adapters.ICart {
	cartModel := CartModel{}
	productModels := []ProductModel{}

	cartResult, cartErr := g.Database.From("carts").Where(goqu.C("checked_out_at").Eq(nil)).ScanStruct(&cartModel)

	if cartErr != nil || !cartResult {
		return &entities.NotFoundCart{}
	}

	productErr := g.Database.Select("items.amount", "p.id", "p.name", "p.price", "p.image", "p.description").
		From("items").
		Join(goqu.T("products").As("p"), goqu.On(goqu.Ex{"items.product_id": goqu.I("p.id")})).
		Where(goqu.C("cart_id").Eq(cartModel.Id)).
		ScanStructs(&productModels)

	if productErr != nil {
		return &entities.NotFoundCart{}
	}

	cart := entities.Cart{}

	for _, model := range productModels {
		for i := 0; i < model.Amount; i++ {
			cart.AddItem(&entities.Product{
				Id:          model.Id,
				Name:        model.Name,
				Price:       model.Price,
				Image:       model.Image,
				Description: model.Description,
			})
		}
	}

	return &cart
}

func (g *CartGateway) SaveCart(cart adapters.ICart) bool {
	return true
}

func (g *CartGateway) UpdateCart(cart adapters.ICart) bool {
	return true
}
