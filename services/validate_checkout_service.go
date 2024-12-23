package services

import (
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/tools/log"
)

func ValidateCheckout(cart *cart.Cart, token string, deps ...interface{}) error {
	for _, a := range cart.Articles {
		err := callValidate(a, token, deps...)
		if err != nil {
			log.Get(deps...).Error(err)
			return err
		}
	}

	return nil
}
