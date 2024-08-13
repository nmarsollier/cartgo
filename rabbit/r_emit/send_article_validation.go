package r_emit

import (
	"encoding/json"

	"github.com/golang/glog"
)

// @Summary		Emite Validar Artículos a Cart cart/article-exist
// @Description	Antes de iniciar las operaciones se validan los artículos contra el catalogo.
// @Tags			Rabbit
// @Accept			json
// @Produce		json
// @Param			body	body	SendValidationMessage	true	"Mensage de validacion"
// @Router			/rabbit/cart/article-exist [put]
//
// Emite Validar Artículos a Cart
func SendArticleValidation(data ArticleValidationData, ctx ...interface{}) error {

	send := SendValidationMessage{
		Type:     "article-exist",
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
		"catalog", // exchange
		"catalog", // routing key
		body,
	)
	if err != nil {
		glog.Error(err)
		chn = nil
		return err
	}

	glog.Info("Rabbit article validation enviado ", string(body))

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