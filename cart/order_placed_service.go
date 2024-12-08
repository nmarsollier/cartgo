package cart

type OrderPlacedEvent struct {
	CartId  string `json:"cartId" example:"CartId"`
	OrderId string `json:"orderId" example:"OrderId"`
	Valid   bool   `json:"valid" example:"true"`
}

func ProcessOrderPlaced(data *OrderPlacedEvent, deps ...interface{}) error {
	cart, err := findById(data.CartId, deps...)
	if err != nil {
		return err
	}

	cart.OrderId = data.OrderId
	err = save(cart, deps...)
	if err != nil {
		return err
	}

	return nil
}
