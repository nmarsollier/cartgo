package cart

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
