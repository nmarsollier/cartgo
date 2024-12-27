package resolvers

import (
	"context"

	"github.com/nmarsollier/cartgo/internal/graph/model"
	"github.com/nmarsollier/cartgo/internal/graph/tools"
)

func CurrentCartResolver(ctx context.Context) (*model.Cart, error) {
	env := tools.GqlDi(ctx)

	user, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return nil, err
	}

	cart, err := env.CartService().CurrentCart(user.ID)
	if err != nil {
		return nil, err
	}

	return cartToModel(cart), nil
}
