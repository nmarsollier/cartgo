package services

import (
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/tools/log"
)

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
