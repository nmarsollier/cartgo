package service

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/tools/apperr"
	"github.com/nmarsollier/cartgo/tools/env"
	"github.com/nmarsollier/cartgo/tools/http_client"
)

func callValidate(article *cart.Article, token string, ctx ...interface{}) error {
	// Buscamos el usuario remoto
	req, err := http.NewRequest("GET", env.Get().CatalogServerURL+"/v1/articles/"+article.ArticleId, nil)
	if err != nil {
		glog.Error(err)
		return apperr.Invalid
	}
	req.Header.Add("Authorization", "bearer "+token)
	resp, err := http_client.Get(ctx...).Do(req)
	if err != nil || resp.StatusCode != 200 {
		glog.Error(err)
		return apperr.Invalid
	}
	defer resp.Body.Close()

	return nil
}
