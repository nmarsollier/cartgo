package services

import (
	"github.com/nmarsollier/cartgo/cart"
)

func FindCartById(cartId string, ctx ...interface{}) (*CartData, error) {
	cart, err := cart.FindCartById(cartId, ctx...)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return nil, err
		}
	}

	return &CartData{
		Id:       cart.ID.Hex(),
		UserId:   cart.UserId,
		OrderId:  cart.OrderId,
		Articles: cart.Articles,
		Enabled:  cart.Enabled,
	}, nil
}
