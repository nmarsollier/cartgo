package resolvers

import (
	"context"

	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/graph/tools"
	"github.com/nmarsollier/cartgo/services"
)

func AddArticle(ctx context.Context, data cart.AddArticleData) (bool, error) {
	user, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return false, err
	}

	env := tools.GqlCtx(ctx)

	_, err = services.AddArticle(user.ID, data, env...)
	if err != nil {
		return false, err
	}

	return true, nil
}
