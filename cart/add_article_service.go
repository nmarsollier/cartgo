package cart

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
