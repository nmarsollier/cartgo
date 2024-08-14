package r_emit

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/nmarsollier/cartgo/cart"
)

//	@Summary		Emite Placed Order desde Cart
//	@Description	Cuando se hace checkout enviamos un comando a orders para que inicie el proceso de la orden.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			body	body	SendPlacedMessage	true	"Place order"
//	@Router			/rabbit/cart/place-order [put]
//
// Emite Placed Order desde Cart
func SendPlaceOrder(cart *cart.Cart, ctx ...interface{}) error {
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

	chn, err := getChannel(ctx...)
	if err != nil {
		glog.Error(err)
		chn = nil
		return err
	}

	err = chn.ExchangeDeclare(
		"order",  // name
		"direct", // type
	)
	if err != nil {
		glog.Error(err)
		chn = nil
		return err
	}

	body, err := json.Marshal(send)
	if err != nil {
		glog.Error(err)
		return err
	}

	err = chn.Publish(
		"order", // exchange
		"order", // routing key
		body,
	)
	if err != nil {
		glog.Error(err)
		chn = nil
		return err
	}

	glog.Info("Rabbit place order enviado ", string(body))
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
	Type     string     `json:"type" example:"place-order"`
	Exchange string     `json:"exchange" example:"order"`
	Queue    string     `json:"queue" example:"order"`
	Message  PlacedData `json:"message" example:"order"`
}
