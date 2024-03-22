package model

type Category struct {
	Id   int      `json:"-" db:"id"`
	Name []string `json:"name" binding:"required"`
}
