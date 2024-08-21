package service

import (
	"net/http"

	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/log"
	"github.com/nmarsollier/cartgo/tools/env"
	"github.com/nmarsollier/cartgo/tools/errs"
	"github.com/nmarsollier/cartgo/tools/httpx"
)

func callValidate(article *cart.Article, token string, ctx ...interface{}) error {
	// Buscamos el usuario remoto
	req, err := http.NewRequest("GET", env.Get().CatalogServerURL+"/v1/articles/"+article.ArticleId, nil)
	if err != nil {
		log.Get(ctx...).Error(err)
		return errs.Invalid
	}
	req.Header.Add("Authorization", "bearer "+token)
	resp, err := httpx.Get(ctx...).Do(req)
	if err != nil || resp.StatusCode != 200 {
		log.Get(ctx...).Error(err)
		return errs.Invalid
	}
	defer resp.Body.Close()

	return nil
}
