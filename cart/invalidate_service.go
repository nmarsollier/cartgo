package cart

func InvalidateCurrentCart(cart *Cart, deps ...interface{}) (err error) {
	return invalidate(cart, deps...)
}
