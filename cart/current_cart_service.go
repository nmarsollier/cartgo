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
