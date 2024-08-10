package r_emit

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

// Emite Validar Artículos a Cart
//
//	@Summary		Emite Validar Artículos a Cart cart/article-exist
//	@Description	Antes de iniciar las operaciones se validan los artículos contra el catalogo.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			body	body	SendValidationMessage	true	"Mensage de validacion"
//
//	@Router			/rabbit/cart/article-exist [put]
func SendArticleValidation(data ArticleValidationData) error {

	send := SendValidationMessage{
		Type:     "article-exist",
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
		"catalog", // name
		"direct",  // type
		false,     // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
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
		"catalog", // exchange
		"catalog", // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			Body: []byte(body),
		})
	if err != nil {
		chn = nil
		return err
	}

	log.Output(1, "Rabbit article validation enviado ")
	return nil
}

type ArticleValidationData struct {
	ReferenceId string `json:"referenceId"`

	ArticleId string `json:"articleId"`
}

type SendValidationMessage struct {
	Type     string                `json:"type"`
	Exchange string                `json:"exchange"`
	Queue    string                `json:"queue"`
	Message  ArticleValidationData `json:"message"`
}
