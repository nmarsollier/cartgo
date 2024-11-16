package cart

func AddArticle(userId string, articleId string, quantity int, ctx ...interface{}) (*Cart, error) {
	cart, err := CurrentCart(userId, ctx...)
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

	cart, err = replace(cart, ctx...)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
