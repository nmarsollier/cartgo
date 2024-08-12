package service

import "github.com/nmarsollier/cartgo/cart"

type ServiceDao interface {
	CallValidate(article *cart.Article, token string) error
}

func Get(options ...interface{}) ServiceDao {
	for _, o := range options {
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
