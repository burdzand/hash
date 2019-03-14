package models

type Product struct {
	ID             int
	Price_in_cents int32
}

func (Product) TableName() string {
	return "app_product"
}
