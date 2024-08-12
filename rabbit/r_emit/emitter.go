package r_emit

import "github.com/nmarsollier/cartgo/cart"

type RabbitEmiter interface {
	SendArticleValidation(data ArticleValidationData) error
	SendPlaceOrder(cart *cart.Cart) error
}

func Get(options ...interface{}) RabbitEmiter {
	for _, o := range options {
		if ti, ok := o.(RabbitEmiter); ok {
			return ti
		}
	}

	return &rabbitEmiter{}
}

type rabbitEmiter struct {
}

func (m *rabbitEmiter) SendArticleValidation(data ArticleValidationData) error {
	return sendArticleValidation(data)
}

func (m *rabbitEmiter) SendPlaceOrder(cart *cart.Cart) error {
	return sendPlaceOrder(cart)
}
