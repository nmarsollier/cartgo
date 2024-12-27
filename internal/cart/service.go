package cart

import "github.com/nmarsollier/cartgo/internal/engine/log"

type CartService interface {
	CurrentCart(userId string) (*Cart, error)
	AddArticle(userId string, articleId string, quantity int) (*Cart, error)
	RemoveArticle(userId string, articleId string) (*Cart, error)
	ProcessOrderPlaced(data *OrderPlacedEvent) error
	InvalidateCurrentCart(cry *Cart) (*Cart, error)
	FindCartById(cartId string) (*Cart, error)
	ProcessArticleData(data *ValidationEvent) error
}

func NewCartService(log log.LogRusEntry, repository CartRepository) CartService {
	return &cartService{
		log:        log,
		repository: repository,
	}
}

type cartService struct {
	log        log.LogRusEntry
	repository CartRepository
}

func (s *cartService) RemoveArticle(userId string, articleId string) (*Cart, error) {
	cart, err := s.CurrentCart(userId)
	if err != nil {
		return nil, err
	}

	newArticles := []*Article{}
	for _, a := range cart.Articles {
		if a.ArticleId != articleId {
			newArticles = append(newArticles, a)
		}
	}
	cart.Articles = newArticles

	cart, err = s.repository.Replace(cart)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

type OrderPlacedEvent struct {
	CartId  string `json:"cartId" example:"CartId"`
	OrderId string `json:"orderId" example:"OrderId"`
	Valid   bool   `json:"valid" example:"true"`
}

func (s *cartService) ProcessOrderPlaced(data *OrderPlacedEvent) error {
	cart, err := s.repository.FindById(data.CartId)
	if err != nil {
		return err
	}

	cart.OrderId = data.OrderId
	_, err = s.repository.Replace(cart)
	if err != nil {
		return err
	}

	return nil
}

func (s *cartService) InvalidateCurrentCart(cart *Cart) (*Cart, error) {
	cart, err := s.repository.Invalidate(cart)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (s *cartService) FindCartById(cartId string) (*Cart, error) {
	cart, err := s.repository.FindById(cartId)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return nil, err
		}
	}

	return cart, nil
}

func (s *cartService) CurrentCart(userId string) (*Cart, error) {
	cart, err := s.repository.FindByUserId(userId)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return nil, err

		}

		cart = s.repository.NewCart(userId)
		cart, err = s.repository.Insert(cart)
		if err != nil {
			return nil, err
		}
	}

	return cart, nil
}

type ValidationEvent struct {
	ReferenceId string `json:"referenceId" example:"UserId"`
	ArticleId   string `json:"articleId" example:"ArticleId"`
	Valid       bool   `json:"valid" example:"true"`
}

func (s *cartService) ProcessArticleData(data *ValidationEvent) error {
	cart, err := s.repository.FindByUserId(data.ReferenceId)
	if err != nil {
		return err
	}

	for _, a := range cart.Articles {
		if a.ArticleId == data.ArticleId {
			a.Validated = true
			a.Valid = data.Valid
		}
	}

	_, err = s.repository.Replace(cart)
	if err != nil {
		return err
	}

	return nil
}

func (s *cartService) AddArticle(userId string, articleId string, quantity int) (*Cart, error) {
	cart, err := s.CurrentCart(userId)
	if err != nil {
		return nil, err
	}

	article := &Article{
		ArticleId: articleId,
		Quantity:  quantity,
		Validated: false,
	}

	// Si existe solo incremento la cantidad
	exist := false
	for _, a := range cart.Articles {
		if a.ArticleId == article.ArticleId {
			a.Quantity += article.Quantity
			exist = true
		}
	}

	// Sino lo agregamos a la lista
	if !exist {
		cart.Articles = append(cart.Articles, article)
	}

	newArticles := []*Article{}
	for _, a := range cart.Articles {
		if a.Quantity > 0 {
			newArticles = append(newArticles, a)
		}
	}
	cart.Articles = newArticles

	cart, err = s.repository.Replace(cart)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
