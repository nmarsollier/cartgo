package resolvers

import (
	"context"

	"github.com/nmarsollier/cartgo/internal/graph/tools"
)

func ValidateCartResolver(ctx context.Context) (bool, error) {
	user, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return false, err
	}
	token, err := tools.TokenString(ctx)

	env := tools.GqlDi(ctx)

	currentCart, err := env.CartService().CurrentCart(user.ID)
	if err != nil {
		return false, err

	}

	err = env.Service().ValidateCheckout(currentCart, token)
	if err != nil {
		return false, err

	}

	return true, nil
}
