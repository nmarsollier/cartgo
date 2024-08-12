package cart

func CurrentCart(userId string, options ...interface{}) (*Cart, error) {
	cart, err := findByUserId(userId, options...)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return nil, err

		}

		cart = newCart(userId)
		cart, err = insert(cart, options...)
		if err != nil {
			return nil, err
		}
	}

	return cart, nil
}

type AddArticleData struct {
	ArticleId string `bson:"articleId" validate:"required,min=1,max=100"`
	Quantity  int    `bson:"quantity" validate:"required,min=1,max=100"`
}

func AddArticle(userId string, articleData AddArticleData, options ...interface{}) (*Cart, error) {
	cart, err := CurrentCart(userId, options...)
	if err != nil {
		return nil, err
	}

	article := &Article{
		ArticleId: articleData.ArticleId,
		Quantity:  articleData.Quantity,
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

	cart, err = replace(cart, options...)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func RemoveArticle(userId string, articleId string, options ...interface{}) (*Cart, error) {
	cart, err := CurrentCart(userId, options...)
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

	cart, err = replace(cart, options...)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func InvalidateCurrentCart(cart *Cart, options ...interface{}) (*Cart, error) {
	cart, err := invalidate(cart, options...)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

type ValidationEvent struct {
	ReferenceId string `json:"referenceId"`
	ArticleId   string `json:"articleId"`
	Valid       bool   `json:"valid"`
}

func ProcessArticleData(data *ValidationEvent, options ...interface{}) error {
	cart, err := findByUserId(data.ReferenceId, options...)
	if err != nil {
		return err
	}

	for _, a := range cart.Articles {
		if a.ArticleId == data.ArticleId {
			a.Validated = true
			a.Valid = data.Valid
		}
	}

	_, err = replace(cart, options...)
	if err != nil {
		return err
	}

	return nil
}

type OrderPlacedEvent struct {
	CartId  string `json:"cartId"`
	OrderId string `json:"orderId"`
	Valid   bool   `json:"valid"`
}

func ProcessOrderPlaced(data *OrderPlacedEvent, options ...interface{}) error {
	cart, err := findById(data.CartId, options...)
	if err != nil {
		return err
	}

	cart.OrderId = data.OrderId
	_, err = replace(cart, options...)
	if err != nil {
		return err
	}

	return nil
}
