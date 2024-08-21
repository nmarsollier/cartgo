package emit

import (
	"encoding/json"

	"github.com/nmarsollier/cartgo/log"
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
	logger := log.Get(ctx...).
		WithField(log.LOG_FIELD_CONTOROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Emit").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "article_exist").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "article_exist")

	corrId, _ := logger.Data[log.LOG_FIELD_CORRELATION_ID].(string)
	send := SendValidationMessage{
		CorrelationId: corrId,
		Exchange:      "article_exist",
		RoutingKey:    "cart_article_exist",
		Message:       data,
	}

	chn, err := getChannel(ctx...)
	if err != nil {
		logger.Error(err)
		chn = nil
		return err
	}

	err = chn.ExchangeDeclare(
		"catalog", // name
		"direct",  // type
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
		"article_exist", // exchange
		"article_exist", // routing key
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

type ArticleValidationData struct {
	ReferenceId string `json:"referenceId" example:"UserId"`

	ArticleId string `json:"articleId" example:"ArticleId"`
}

type SendValidationMessage struct {
	CorrelationId string                `json:"correlation_id" example:"123123" `
	Exchange      string                `json:"exchange" example:"cart"`
	RoutingKey    string                `json:"routing_key" example:""`
	Message       ArticleValidationData `json:"message"`
}
