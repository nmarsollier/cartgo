package services

import (
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/rabbit/emit"
)

func AddArticle(userId string, articleID string, quantity int, deps ...interface{}) (*cart.Cart, error) {
	cart, err := cart.AddArticle(userId, articleID, quantity, deps...)
	if err != nil {
		return nil, err
	}

	for _, a := range cart.Articles {
		if !a.Validated {
			emit.SendArticleValidation(
				emit.ArticleValidationData{
					ReferenceId: cart.UserId,
					ArticleId:   a.ArticleId,
				},
				deps...)
		}
	}

	return cart, nil
}
