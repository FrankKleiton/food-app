package entities

import "food-app/domain/adapters"

type NotFoundCart struct{}

func (c *NotFoundCart) AddItem(p adapters.IProduct) {}

func (c *NotFoundCart) GetItem(id string) adapters.IItem {
	return nil
}

func (c *NotFoundCart) GetTotal() float64 {
	return -1111.11
}

func (c *NotFoundCart) GetItems() []adapters.IItem {
	return nil
}

type Cart struct {
	Items []adapters.IItem
}

func (c *Cart) AddItem(p adapters.IProduct) {
	if cp := c.GetItem(p.GetId()); cp != nil {
		cp.Increment(1)
	} else {
		c.Items = append(c.Items, &Item{Amount: 1, Product: p})
	}
}

func (c *Cart) GetItem(id string) adapters.IItem {
	for _, p := range c.Items {
		if p.GetId() == id {
			return p
		}
	}

	return nil
}

func (c *Cart) GetTotal() float64 {
	var total float64

	for _, p := range c.Items {
		total += p.GetTotal()
	}

	return total
}

func (c *Cart) GetItems() []adapters.IItem {
	return c.Items
}
