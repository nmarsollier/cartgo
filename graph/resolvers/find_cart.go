package resolvers

import (
	"context"

	"github.com/nmarsollier/cartgo/graph/tools"
	"github.com/nmarsollier/cartgo/services"
)

func FindCartById(ctx context.Context, id string) (*services.CartData, error) {
	env := tools.GqlCtx(ctx)

	/*	_, err := tools.ValidateLoggedIn(ctx)
		if err != nil {
			return nil, err
		}*/

	cart, err := services.FindCartById(id, env...)

	if err != nil {
		return nil, err
	}

	return cart, nil
}
