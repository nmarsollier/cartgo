package services

import (
	"net/http"

	"github.com/nmarsollier/cartgo/internal/cart"
	"github.com/nmarsollier/cartgo/internal/engine/env"
	"github.com/nmarsollier/cartgo/internal/engine/errs"
	"github.com/nmarsollier/cartgo/internal/engine/httpx"
	"github.com/nmarsollier/cartgo/internal/engine/log"
	"github.com/nmarsollier/cartgo/internal/rabbit/emit"
)

type Service interface {
	AddArticle(userId string, articleID string, quantity int) (*cart.Cart, error)
	Checkout(userId string, token string) (*cart.Cart, error)
	GetCurrentCart(userId string) (*cart.Cart, error)
	FindCartById(cartId string) (*cart.Cart, error)
	ValidateCheckout(cart *cart.Cart, token string) error
}

func NewService(log log.LogRusEntry, http httpx.HTTPClient, cart cart.CartService, emit emit.RabbitEmiter) Service {
	return &service{
		log:  log,
		http: http,
		cart: cart,
		emit: emit,
	}
}

type service struct {
	log  log.LogRusEntry
	http httpx.HTTPClient
	cart cart.CartService
	emit emit.RabbitEmiter
}

func (s *service) AddArticle(userId string, articleID string, quantity int) (*cart.Cart, error) {
	cart, err := s.cart.AddArticle(userId, articleID, quantity)
	if err != nil {
		return nil, err
	}

	for _, a := range cart.Articles {
		if !a.Validated {
			s.emit.SendArticleValidation(
				emit.ArticleValidationData{
					ReferenceId: cart.UserId,
					ArticleId:   a.ArticleId,
				})
		}
	}

	return cart, nil
}

func (s *service) Checkout(userId string, token string) (*cart.Cart, error) {
	currentCart, err := s.cart.CurrentCart(userId)
	if err != nil {
		return nil, err
	}

	err = s.ValidateCheckout(currentCart, token)
	if err != nil {
		return nil, err
	}

	currentCart, err = s.cart.InvalidateCurrentCart(currentCart)
	if err != nil {
		return nil, err
	}

	s.emit.SendPlaceOrder(currentCart)

	return currentCart, nil
}

func (s *service) callValidate(article *cart.Article, token string) error {
	// Buscamos el usuario remoto
	req, err := http.NewRequest("GET", env.Get().CatalogServerURL+"/articles/"+article.ArticleId, nil)
	if corrId, ok := s.log.Data()[log.LOG_FIELD_CORRELATION_ID].(string); ok {
		req.Header.Add(log.LOG_FIELD_CORRELATION_ID, corrId)
	}

	if err != nil {
		s.log.Error(err)
		return errs.Invalid
	}
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := s.http.Do(req)
	if err != nil || resp.StatusCode != 200 {
		s.log.Error(err)
		return errs.Invalid
	}
	defer resp.Body.Close()

	return nil
}

func (s *service) GetCurrentCart(userId string) (*cart.Cart, error) {
	cart, err := s.cart.CurrentCart(userId)
	if err != nil {
		return nil, err
	}

	for _, a := range cart.Articles {
		if !a.Validated {
			s.emit.SendArticleValidation(
				emit.ArticleValidationData{
					ReferenceId: cart.UserId,
					ArticleId:   a.ArticleId,
				},
			)
		}
	}

	return cart, nil
}
func (s *service) FindCartById(cartId string) (*cart.Cart, error) {
	cart, err := s.cart.FindCartById(cartId)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return nil, err
		}
	}

	return cart, nil
}
func (s *service) ValidateCheckout(cart *cart.Cart, token string) error {
	for _, a := range cart.Articles {
		err := s.callValidate(a, token)
		if err != nil {
			s.log.Error(err)
			return err
		}
	}

	return nil
}
