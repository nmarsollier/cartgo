package resolvers

import (
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/graph/model"
)

func cartToModel(cart *cart.Cart) *model.Cart {
	return &model.Cart{
		ID:       cart.ID.Hex(),
		UserID:   cart.UserId,
		OrderID:  &cart.OrderId,
		Articles: articlesToModel(cart.Articles),
		Enabled:  cart.Enabled,
	}
}

func articlesToModel(articles []*cart.Article) []*model.CartArticle {
	result := make([]*model.CartArticle, len(articles))
	for i, a := range articles {
		result[i] = &model.CartArticle{
			ArticleID: a.ArticleId,
			Article:   &model.Article{ID: a.ArticleId},
			Quantity:  a.Quantity,
			Valid:     a.Valid,
			Validated: a.Validated,
		}
	}
	return result
}
