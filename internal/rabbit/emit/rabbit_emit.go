package emit

import (
	"github.com/nmarsollier/commongo/rbt"
)

type ArticleValidationPublisher = rbt.RabbitPublisher[*ArticleValidationData]
type PlacedDataPublisher = rbt.RabbitPublisher[*PlacedData]

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
