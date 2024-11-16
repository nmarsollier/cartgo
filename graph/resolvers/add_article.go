package resolvers

import (
	"context"

	"github.com/nmarsollier/cartgo/graph/tools"
	"github.com/nmarsollier/cartgo/services"
)

func AddArticle(ctx context.Context, articleID string, quantity int) (bool, error) {
	user, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return false, err
	}

	env := tools.GqlCtx(ctx)

	_, err = services.AddArticle(user.ID, articleID, quantity, env...)
	if err != nil {
		return false, err
	}

	return true, nil
}
