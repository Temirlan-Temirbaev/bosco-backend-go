package model

type CategoryItem struct {
	Id         int
	ProductId  int `json:"product_id" db:"product_id"`
	CategoryId int `json:"category_id" db:"product_id"`
}
