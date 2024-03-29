package model

type Contact struct {
	Id          int    `json:"-" db:"id"`
	Address     string `json:"address" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	VipPhone    string `json:"vip_phone" binding:"required" db:"vip_phone"`
	Coordinates string `json:"coordinates" binding:"required"`
}
