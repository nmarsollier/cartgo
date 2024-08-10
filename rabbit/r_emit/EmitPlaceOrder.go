package r_emit

import (
	"encoding/json"
	"log"

	"github.com/nmarsollier/cartgo/cart"
	"github.com/streadway/amqp"
)

// Emite Placed Order desde Cart
//
//	@Summary		Emite Placed Order desde Cart
//	@Description	Emite Placed Order desde Cart
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			body	body	SendPlacedMessage	true	"Mensage de validacion"
//
//	@Router			/rabbit/cart/place-order [put]
func SendPlaceOrder(cart *cart.Cart) error {
	articles := []PlaceArticlesData{}
	for _, a := range cart.Articles {
		articles = append(articles, PlaceArticlesData{
			a.ArticleId,
			a.Quantity,
		})
	}

	data := PlacedData{
		CartId:   cart.ID.Hex(),
		UserId:   cart.UserId,
		Articles: articles,
	}

	send := SendPlacedMessage{
		Type:     "place-order",
		Exchange: "cart",
		Queue:    "cart",
		Message:  data,
	}

	chn, err := getChannel()
	if err != nil {
		chn = nil
		return err
	}

	err = chn.ExchangeDeclare(
		"order",  // name
		"direct", // type
		false,    // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		chn = nil
		return err
	}

	body, err := json.Marshal(send)
	if err != nil {
		return err
	}

	err = chn.Publish(
		"order", // exchange
		"order", // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			Body: []byte(body),
		})
	if err != nil {
		chn = nil
		return err
	}

	log.Output(1, "Rabbit place order enviado ")
	return nil
}

type PlacedData struct {
	CartId   string              `json:"cartId"`
	UserId   string              `json:"userId"`
	Articles []PlaceArticlesData `json:"articles"`
}

type PlaceArticlesData struct {
	Id       string `json:"id"`
	Quantity int    `json:"quantity"`
}

type SendPlacedMessage struct {
	Type     string     `json:"type"`
	Exchange string     `json:"exchange"`
	Queue    string     `json:"queue"`
	Message  PlacedData `json:"message"`
}
