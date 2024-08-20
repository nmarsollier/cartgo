package emit

import (
	"encoding/json"

	"github.com/golang/glog"
)

//	@Summary		Emite Validar Artículos a Cart article_exist/article_exist
//	@Description	Solicitamos las validaciones ar articulos a catalogo. Responde en article_exist/cart_article_exist.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			body	body	SendValidationMessage	true	"Mensage de validacion article_exist/cart_article_exist"
//	@Router			/rabbit/article_exist [put]
//
// Emite Validar Artículos a Cart
func SendArticleValidation(data ArticleValidationData, ctx ...interface{}) error {

	send := SendValidationMessage{
		Exchange:   "article_exist",
		RoutingKey: "cart_article_exist",
		Message:    data,
	}

	chn, err := getChannel(ctx...)
	if err != nil {
		glog.Error(err)
		chn = nil
		return err
	}

	err = chn.ExchangeDeclare(
		"catalog", // name
		"direct",  // type
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
		"article_exist", // exchange
		"article_exist", // routing key
		body,
	)
	if err != nil {
		glog.Error(err)
		chn = nil
		return err
	}

	glog.Info("Emit article_exist :", string(body))

	return nil
}

type ArticleValidationData struct {
	ReferenceId string `json:"referenceId" example:"UserId"`

	ArticleId string `json:"articleId" example:"ArticleId"`
}

type SendValidationMessage struct {
	Exchange   string                `json:"exchange" example:"cart"`
	RoutingKey string                `json:"routing_key" example:""`
	Message    ArticleValidationData `json:"message"`
}
