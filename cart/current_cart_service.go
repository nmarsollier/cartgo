package cart

func CurrentCart(userId string, deps ...interface{}) (*Cart, error) {
	cart, err := findByUserId(userId, deps...)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return nil, err

		}

		cart = newCart(userId)
		cart, err = insert(cart, deps...)
		if err != nil {
			return nil, err
		}
	}

	return cart, nil
}
