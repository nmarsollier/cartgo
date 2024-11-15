package resolvers

import (
	"context"

	"github.com/nmarsollier/cartgo/graph/tools"
	"github.com/nmarsollier/cartgo/services"
)

func CheckoutResolver(ctx context.Context) (bool, error) {
	user, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return false, err
	}
	tokenString, err := tools.TokenString(ctx)

	env := tools.GqlCtx(ctx)

	_, err = services.Checkout(user.ID, tokenString, env...)
	if err != nil {
		return false, err
	}
	return true, nil
}
