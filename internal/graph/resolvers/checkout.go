package resolvers

import (
	"context"

	"github.com/nmarsollier/cartgo/internal/graph/tools"
)

func CheckoutResolver(ctx context.Context) (bool, error) {
	user, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return false, err
	}
	tokenString, err := tools.TokenString(ctx)

	env := tools.GqlDi(ctx)

	_, err = env.Service().Checkout(user.ID, tokenString)
	if err != nil {
		return false, err
	}
	return true, nil
}
