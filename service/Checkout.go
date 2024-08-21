package service

import (
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/log"
	"github.com/nmarsollier/cartgo/rabbit/emit"
)

func Checkout(userId string, token string, ctx ...interface{}) (*cart.Cart, error) {
	currentCart, err := cart.CurrentCart(userId, ctx...)
	if err != nil {
		return nil, err
	}

	err = ValidateCheckout(currentCart, token, ctx...)
	if err != nil {
		return nil, err
	}

	currentCart, err = cart.InvalidateCurrentCart(currentCart, ctx...)
	if err != nil {
		return nil, err
	}

	emit.SendPlaceOrder(currentCart, ctx...)

	return currentCart, nil
}

func ValidateCheckout(cart *cart.Cart, token string, ctx ...interface{}) error {
	for _, a := range cart.Articles {
		err := callValidate(a, token, ctx...)
		if err != nil {
			log.Get(ctx...).Error(err)
			return err
		}
	}

	return nil
}
