package service

import (
	"net/http"

	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/rabbit/r_emit"
	"github.com/nmarsollier/cartgo/tools/env"
	"github.com/nmarsollier/cartgo/tools/errors"
)

func Checkout(userId string, token string) (*cart.Cart, error) {
	currentCart, err := cart.CurrentCart(userId)
	if err != nil {
		return nil, err
	}

	err = ValidateCheckout(currentCart, token)
	if err != nil {
		return nil, err
	}

	currentCart, err = cart.InvalidateCurrentCart(userId)
	if err != nil {
		return nil, err
	}

	r_emit.SendPlaceOrder(currentCart)

	return currentCart, nil
}

func ValidateCheckout(cart *cart.Cart, token string) error {
	for _, a := range cart.Articles {
		err := callValidate(a, token)
		if err != nil {
			return err
		}
	}

	return nil
}

func callValidate(article *cart.Article, token string) error {
	// Buscamos el usuario remoto
	req, err := http.NewRequest("GET", env.Get().CatalogServerURL+"/v1/articles/"+article.ArticleId, nil)
	if err != nil {
		return errors.Unauthorized
	}
	req.Header.Add("Authorization", "bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return errors.Unauthorized
	}
	defer resp.Body.Close()

	return nil
}
