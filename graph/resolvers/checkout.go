package resolvers

import (
	"context"

	"github.com/nmarsollier/cartgo/graph/tools"
	"github.com/nmarsollier/cartgo/services"
)

func CheckoutResolver(ctx context.Context) (result bool, err error) {
	user, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return
	}
	tokenString, err := tools.TokenString(ctx)
	if err != nil {
		return
	}

	env := tools.GqlCtx(ctx)

	err = services.Checkout(user.ID, tokenString, env...)
	result = err == nil
	return
}
