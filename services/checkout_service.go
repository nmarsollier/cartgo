package services

import (
	"net/http"

	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/rabbit/emit"
	"github.com/nmarsollier/cartgo/tools/log"

	"github.com/nmarsollier/cartgo/tools/env"
	"github.com/nmarsollier/cartgo/tools/errs"
	"github.com/nmarsollier/cartgo/tools/httpx"
)

func Checkout(userId string, token string, deps ...interface{}) (err error) {
	currentCart, err := cart.CurrentCart(userId, deps...)
	if err != nil {
		return
	}

	err = ValidateCheckout(currentCart, token, deps...)
	if err != nil {
		return
	}

	err = cart.InvalidateCurrentCart(currentCart, deps...)
	if err != nil {
		return
	}

	emit.SendPlaceOrder(currentCart, deps...)

	return
}

func callValidate(article *cart.Article, token string, deps ...interface{}) error {
	// Buscamos el usuario remoto
	req, err := http.NewRequest("GET", env.Get().CatalogServerURL+"/v1/articles/"+article.ArticleId, nil)
	if corrId, ok := log.Get(deps...).Data()[log.LOG_FIELD_CORRELATION_ID].(string); ok {
		req.Header.Add(log.LOG_FIELD_CORRELATION_ID, corrId)
	}

	if err != nil {
		log.Get(deps...).Error(err)
		return errs.Invalid
	}
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := httpx.Get(deps...).Do(req)
	if err != nil || resp.StatusCode != 200 {
		log.Get(deps...).Error(err)
		return errs.Invalid
	}
	defer resp.Body.Close()

	return nil
}
