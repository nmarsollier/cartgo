package resolvers

import (
	"context"

	"github.com/nmarsollier/cartgo/internal/graph/tools"
)

func IncrementArticleResolver(ctx context.Context, articleID string) (bool, error) {
	user, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return false, err
	}

	env := tools.GqlDi(ctx)

	_, err = env.Service().AddArticle(user.ID, articleID, 1)
	if err != nil {
		return false, err
	}

	return true, nil
}
