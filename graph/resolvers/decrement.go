package resolvers

import (
	"context"

	"github.com/nmarsollier/cartgo/graph/tools"
	"github.com/nmarsollier/cartgo/services"
)

func DecrementArticleResolver(ctx context.Context, articleID string) (bool, error) {
	user, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return false, err
	}

	env := tools.GqlCtx(ctx)

	_, err = services.AddArticle(user.ID, articleID, -1, env...)
	if err != nil {
		return false, err
	}

	return true, nil
}
