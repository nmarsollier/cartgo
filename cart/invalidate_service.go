package cart

func InvalidateCurrentCart(cart *Cart, ctx ...interface{}) (*Cart, error) {
	cart, err := invalidate(cart, ctx...)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
