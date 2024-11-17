package resolvers

import (
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/graph/model"
)

func cartToModel(cart *cart.Cart) *model.Cart {
	var Order *model.Order
	if cart.OrderId != "" {
		Order = &model.Order{ID: cart.OrderId}
	}

	var User *model.User
	if cart.UserId != "" {
		User = &model.User{ID: cart.UserId}
	}

	return &model.Cart{
		ID:       cart.ID.Hex(),
		UserID:   cart.UserId,
		User:     User,
		OrderID:  &cart.OrderId,
		Order:    Order,
		Articles: articlesToModel(cart.Articles),
		Enabled:  cart.Enabled,
	}
}

func articlesToModel(articles []*cart.Article) []*model.CartArticle {
	result := make([]*model.CartArticle, len(articles))
	for i, a := range articles {
		var Article *model.Article
		if a.ArticleId != "" && a.Valid {
			Article = &model.Article{ID: a.ArticleId}
		}

		result[i] = &model.CartArticle{
			ArticleID: a.ArticleId,
			Article:   Article,
			Quantity:  a.Quantity,
			Valid:     a.Valid,
			Validated: a.Validated,
		}
	}
	return result
}
