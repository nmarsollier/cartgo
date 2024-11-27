package cart

func InvalidateCurrentCart(cart *Cart, deps ...interface{}) (*Cart, error) {
	cart, err := invalidate(cart, deps...)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
