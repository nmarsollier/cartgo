package resolvers

import (
	"context"

	"github.com/nmarsollier/cartgo/internal/graph/tools"
)

func AddArticle(ctx context.Context, articleID string, quantity int) (bool, error) {
	user, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return false, err
	}

	env := tools.GqlDi(ctx)

	_, err = env.Service().AddArticle(user.ID, articleID, quantity)
	if err != nil {
		return false, err
	}

	return true, nil
}
