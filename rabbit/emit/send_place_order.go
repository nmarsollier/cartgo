package emit

import (
	"encoding/json"

	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/tools/log"
)

//	@Summary		Emite place_order/place_order
//	@Description	Cuando se hace checkout enviamos un comando a orders para que inicie el proceso de la orden.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			body	body	SendPlacedMessage	true	"Place order"
//	@Router			/rabbit/place_order [put]
//
// Emite Placed Order desde Cart
func SendPlaceOrder(cart *cart.Cart, deps ...interface{}) error {
	logger := log.Get(deps...).
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Emit").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "place_order").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "place_order")

	articles := []PlaceArticlesData{}
	for _, a := range cart.Articles {
		articles = append(articles, PlaceArticlesData{
			a.ArticleId,
			a.Quantity,
		})
	}

	data := PlacedData{
		CartId:   cart.ID,
		UserId:   cart.UserId,
		Articles: articles,
	}

	corrId, _ := logger.Data()[log.LOG_FIELD_CORRELATION_ID].(string)
	send := SendPlacedMessage{
		CorrelationId: corrId,
		Message:       data,
	}

	chn, err := getChannel(deps...)
	if err != nil {
		logger.Error(err)
		chn = nil
		return err
	}

	err = chn.ExchangeDeclare(
		"place_order", // name
		"direct",      // type
	)
	if err != nil {
		logger.Error(err)
		chn = nil
		return err
	}

	body, err := json.Marshal(send)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = chn.Publish(
		"place_order", // exchange
		"place_order", // routing key
		body,
	)
	if err != nil {
		logger.Error(err)
		chn = nil
		return err
	}

	logger.Info(string(body))
	return nil
}

type PlacedData struct {
	CartId   string              `json:"cartId" example:"CartId"`
	UserId   string              `json:"userId" example:"UserId"`
	Articles []PlaceArticlesData `json:"articles"`
}

type PlaceArticlesData struct {
	Id       string `json:"id" example:"ArticleId"`
	Quantity int    `json:"quantity" example:"10"`
}

type SendPlacedMessage struct {
	CorrelationId string     `json:"correlation_id" example:"123123" `
	Message       PlacedData `json:"message" example:"order"`
}
