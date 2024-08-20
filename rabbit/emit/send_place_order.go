package emit

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/nmarsollier/cartgo/cart"
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
		Message: data,
	}

	chn, err := getChannel(ctx...)
	if err != nil {
		glog.Error(err)
		chn = nil
		return err
	}

	err = chn.ExchangeDeclare(
		"place_order", // name
		"direct",      // type
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
		"place_order", // exchange
		"place_order", // routing key
		body,
	)
	if err != nil {
		glog.Error(err)
		chn = nil
		return err
	}

	glog.Info("Emit place_order :", string(body))
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
	Message PlacedData `json:"message" example:"order"`
}
