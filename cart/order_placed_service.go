package cart

type OrderPlacedEvent struct {
	CartId  string `json:"cartId" example:"CartId"`
	OrderId string `json:"orderId" example:"OrderId"`
	Valid   bool   `json:"valid" example:"true"`
}

func ProcessOrderPlaced(data *OrderPlacedEvent, ctx ...interface{}) error {
	cart, err := findById(data.CartId, ctx...)
	if err != nil {
		return err
	}

	cart.OrderId = data.OrderId
	_, err = replace(cart, ctx...)
	if err != nil {
		return err
	}

	return nil
}
