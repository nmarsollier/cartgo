package resolvers

import (
	"context"

	"github.com/nmarsollier/cartgo/internal/graph/model"
	"github.com/nmarsollier/cartgo/internal/graph/tools"
)

func FindCartById(ctx context.Context, id string) (*model.Cart, error) {
	env := tools.GqlDi(ctx)

	_, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return nil, err
	}

	cart, err := env.CartService().FindCartById(id)

	if err != nil {
		return nil, err
	}

	return cartToModel(cart), nil
}
