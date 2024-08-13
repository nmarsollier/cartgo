package service

import (
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/rabbit/r_emit"
)

func AddArticle(userId string, article cart.AddArticleData, ctx ...interface{}) (*cart.Cart, error) {
	cart, err := cart.AddArticle(userId, article, ctx...)
	if err != nil {
		return nil, err
	}

	for _, a := range cart.Articles {
		if !a.Validated {
			r_emit.SendArticleValidation(r_emit.ArticleValidationData{
				ReferenceId: cart.UserId,
				ArticleId:   a.ArticleId,
			}, ctx...)
		}
	}

	return cart, nil
}
