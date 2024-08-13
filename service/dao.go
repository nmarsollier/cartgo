package service

import "github.com/nmarsollier/cartgo/cart"

type ServiceDao interface {
	CallValidate(article *cart.Article, token string) error
}

func Get(ctx ...interface{}) ServiceDao {
	for _, o := range ctx {
		if ti, ok := o.(ServiceDao); ok {
			return ti
		}
	}

	return &httpDaoImpl{}
}

type httpDaoImpl struct {
}

func (t *httpDaoImpl) CallValidate(article *cart.Article, token string) error {
	return callValidate(article, token)
}
