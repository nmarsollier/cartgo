package services

import "github.com/nmarsollier/cartgo/cart"

type CartData struct {
	Id       string          `json:"_id"`
	UserId   string          `json:"userId" validate:"required,min=1,max=100"`
	OrderId  string          `json:"orderId"`
	Articles []*cart.Article `json:"articles" validate:"required"`
	Enabled  bool            `json:"enabled"`
}
