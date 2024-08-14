package cart

func CurrentCart(userId string, ctx ...interface{}) (*Cart, error) {
	cart, err := findByUserId(userId, ctx...)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return nil, err

		}

		cart = newCart(userId)
		cart, err = insert(cart, ctx...)
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

func AddArticle(userId string, articleData AddArticleData, ctx ...interface{}) (*Cart, error) {
	cart, err := CurrentCart(userId, ctx...)
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

	cart, err = replace(cart, ctx...)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func RemoveArticle(userId string, articleId string, ctx ...interface{}) (*Cart, error) {
	cart, err := CurrentCart(userId, ctx...)
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

	cart, err = replace(cart, ctx...)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func InvalidateCurrentCart(cart *Cart, ctx ...interface{}) (*Cart, error) {
	cart, err := invalidate(cart, ctx...)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

type ValidationEvent struct {
	ReferenceId string `json:"referenceId" example:"UserId"`
	ArticleId   string `json:"articleId" example:"ArticleId"`
	Valid       bool   `json:"valid" example:"true"`
}

func ProcessArticleData(data *ValidationEvent, ctx ...interface{}) error {
	cart, err := findByUserId(data.ReferenceId, ctx...)
	if err != nil {
		return err
	}

	for _, a := range cart.Articles {
		if a.ArticleId == data.ArticleId {
			a.Validated = true
			a.Valid = data.Valid
		}
	}

	_, err = replace(cart, ctx...)
	if err != nil {
		return err
	}

	return nil
}

type OrderPlacedEvent struct {
	CartId  string `json:"cartId" example:"CartId"`
	OrderId string `json:"orderId" example:"OrderId"`
	Valid   bool   `json:"valid" example:"true"`
}

func ProcessOrderPlaced(data *OrderPlacedEvent, ctx ...interface{}) error {
	cart, err := findById(data.CartId, ctx...)
	if err != nil {
		return err
	}

	cart.OrderId = data.OrderId
	_, err = replace(cart, ctx...)
	if err != nil {
		return err
	}

	return nil
}
