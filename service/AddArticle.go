package service

import (
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/rabbit/emit"
)

func AddArticle(userId string, article cart.AddArticleData, ctx ...interface{}) (*cart.Cart, error) {
	cart, err := cart.AddArticle(userId, article, ctx...)
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
				ctx...)
		}
	}

	return cart, nil
}
