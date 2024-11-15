package resolvers

import (
	"context"

	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/graph/tools"
	"github.com/nmarsollier/cartgo/services"
)

func IncrementArticleResolver(ctx context.Context, articleID string) (bool, error) {
	user, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return false, err
	}

	env := tools.GqlCtx(ctx)

	article := cart.AddArticleData{
		ArticleId: articleID,
		Quantity:  1,
	}

	_, err = services.AddArticle(user.ID, article, env...)
	if err != nil {
		return false, err
	}

	return true, nil
}
