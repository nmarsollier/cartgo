package resolvers

import (
	"context"

	"github.com/nmarsollier/cartgo/internal/graph/tools"
)

func RemoveArticleResolver(ctx context.Context, articleID string) (bool, error) {
	user, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return false, err
	}

	env := tools.GqlDi(ctx)

	_, err = env.CartService().RemoveArticle(user.ID, articleID)
	if err != nil {
		return false, err
	}

	return true, nil
}
