package resolvers

import (
	"context"

	"github.com/nmarsollier/cartgo/graph/model"
	"github.com/nmarsollier/cartgo/graph/tools"
	"github.com/nmarsollier/cartgo/services"
)

func CurrentCartResolver(ctx context.Context) (*model.Cart, error) {
	env := tools.GqlCtx(ctx)

	user, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return nil, err
	}

	cart, err := services.GetCurrentCart(user.ID, env...)
	if err != nil {
		return nil, err
	}

	return cartToModel(cart), nil
}
