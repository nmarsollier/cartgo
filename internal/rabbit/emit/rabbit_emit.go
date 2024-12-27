package emit

import (
	"encoding/json"

	"github.com/nmarsollier/cartgo/internal/cart"
	"github.com/nmarsollier/cartgo/internal/engine/log"
)

type RabbitEmiter interface {
	SendArticleValidation(data ArticleValidationData) error
	SendPlaceOrder(cart *cart.Cart) error
}

func NewRabbitEmiter(logger log.LogRusEntry, channel RabbitChannel) RabbitEmiter {
	return &emitArticleValidation{
		logger:  logger,
		channel: channel,
	}
}

type emitArticleValidation struct {
	logger  log.LogRusEntry
	channel RabbitChannel
}

//	@Summary		Emite Validar Artículos a Cart article_exist/article_exist
//	@Description	Solicitamos las validaciones ar articulos a catalogo. Responde en article_exist/cart_article_exist.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			body	body	SendValidationMessage	true	"Mensage de validacion article_exist/cart_article_exist"
//	@Router			/rabbit/article_exist [put]
//
// Emite Validar Artículos a Cart
func (e *emitArticleValidation) SendArticleValidation(data ArticleValidationData) error {
	e.logger.WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Emit").
		WithField(log.LOG_FIELD_RABBIT_EXCHANGE, "article_exist").
		WithField(log.LOG_FIELD_RABBIT_QUEUE, "article_exist")
	corrId, _ := e.logger.Data()[log.LOG_FIELD_CORRELATION_ID].(string)

	send := SendValidationMessage{
		CorrelationId: corrId,
		Exchange:      "article_exist",
		RoutingKey:    "cart_article_exist",
		Message:       data,
	}

	err := e.channel.ExchangeDeclare(
		"catalog", // name
		"direct",  // type
	)
	if err != nil {
		e.logger.Error(err)
		return err
	}

	body, err := json.Marshal(send)
	if err != nil {
		e.logger.Error(err)
		return err
	}

	err = e.channel.Publish(
		"article_exist", // exchange
		"article_exist", // routing key
		body,
	)
	if err != nil {
		e.logger.Error(err)
		return err
	}

	e.logger.Info(string(body))

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

//	@Summary		Emite place_order/place_order
//	@Description	Cuando se hace checkout enviamos un comando a orders para que inicie el proceso de la orden.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			body	body	SendPlacedMessage	true	"Place order"
//	@Router			/rabbit/place_order [put]
//
// Emite Placed Order desde Cart
func (e *emitArticleValidation) SendPlaceOrder(cart *cart.Cart) error {
	e.logger.
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
		CartId:   cart.ID.Hex(),
		UserId:   cart.UserId,
		Articles: articles,
	}

	corrId, _ := e.logger.Data()[log.LOG_FIELD_CORRELATION_ID].(string)
	send := SendPlacedMessage{
		CorrelationId: corrId,
		Message:       data,
	}

	err := e.channel.ExchangeDeclare(
		"place_order", // name
		"direct",      // type
	)
	if err != nil {
		e.logger.Error(err)
		return err
	}

	body, err := json.Marshal(send)
	if err != nil {
		e.logger.Error(err)
		return err
	}

	err = e.channel.Publish(
		"place_order", // exchange
		"place_order", // routing key
		body,
	)
	if err != nil {
		e.logger.Error(err)
		return err
	}

	e.logger.Info(string(body))
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
