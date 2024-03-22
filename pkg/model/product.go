package model

type Product struct {
	Id          int      `json:"-" db:"id"`
	Name        []string `json:"name" binding:"required"`
	Description []string `json:"description"`
	Price       int      `json:"price" binding:"required"`
	Discount    int      `json:"discount"`
}
