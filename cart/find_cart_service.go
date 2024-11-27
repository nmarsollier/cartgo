package cart

func FindCartById(cartId string, deps ...interface{}) (*Cart, error) {
	cart, err := findById(cartId, deps...)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return nil, err
		}
	}

	return cart, nil
}
