package resolvers

import (
	"context"

	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/graph/tools"
	"github.com/nmarsollier/cartgo/services"
)

func ValidateCartResolver(ctx context.Context) (bool, error) {
	user, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return false, err
	}
	token, err := tools.TokenString(ctx)

	env := tools.GqlCtx(ctx)

	currentCart, err := cart.CurrentCart(user.ID, env...)
	if err != nil {
		return false, err

	}

	err = services.ValidateCheckout(currentCart, token, env...)
	if err != nil {
		return false, err

	}

	return true, nil
}
