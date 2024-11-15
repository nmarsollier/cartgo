package cart

func FindCartById(cartId string, ctx ...interface{}) (*Cart, error) {
	cart, err := findById(cartId, ctx...)
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return nil, err
		}
	}

	return cart, nil
}
