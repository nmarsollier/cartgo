package resolvers

import (
	"context"

	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/graph/tools"
)

func RemoveArticleResolver(ctx context.Context, articleID string) (bool, error) {
	user, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return false, err
	}

	env := tools.GqlCtx(ctx)

	_, err = cart.RemoveArticle(user.ID, articleID, env...)
	if err != nil {
		return false, err
	}

	return true, nil
}
