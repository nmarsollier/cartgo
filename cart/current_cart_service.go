package cart

import (
	"github.com/nmarsollier/cartgo/tools/errs"
)

func CurrentCart(userId string, deps ...interface{}) (*Cart, error) {
	cart, err := findByUserId(userId, deps...)
	if err == errs.NotFound {
		cart = newCart(userId)
		err = save(cart, deps...)
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	return cart, nil
}
