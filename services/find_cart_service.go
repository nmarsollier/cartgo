package services

import (
	"github.com/nmarsollier/cartgo/cart"
)

func FindCartById(cartId string, ctx ...interface{}) (*cart.Cart, error) {
	cart, err := cart.FindCartById(cartId, ctx...)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return nil, err
		}
	}

	return cart, nil
}
