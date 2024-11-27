package cart

func RemoveArticle(userId string, articleId string, deps ...interface{}) (*Cart, error) {
	cart, err := CurrentCart(userId, deps...)
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

	cart, err = replace(cart, deps...)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
