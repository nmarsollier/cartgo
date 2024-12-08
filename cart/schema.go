package cart

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Estuctura basica de del evento
type Cart struct {
	ID            string     `dynamodbav:"id" json:"_id"`
	UserId        string     `dynamodbav:"userId"  json:"userId" validate:"required,min=1,max=100"`
	UserIdEnabled string     `dynamodbav:"userId_enabled"`
	OrderId       string     `dynamodbav:"orderId" json:"orderId"`
	Articles      []*Article `dynamodbav:"articles"  json:"articles" validate:"required"`
	Enabled       bool       `dynamodbav:"enabled" json:"enabled"`
	Created       time.Time  `dynamodbav:"created" json:"created"`
	Updated       time.Time  `dynamodbav:"updated" json:"updated"`
}

// validateSchema valida la estructura para ser insertada en la db
func (e *Cart) validateSchema() error {
	return validator.New().Struct(e)
}

type Article struct {
	ArticleId string `dynamodbav:"articleId" json:"articleId" validate:"required,min=1,max=100"`
	Quantity  int    `dynamodbav:"quantity" json:"quantity" validate:"required,min=1,max=100"`
	Valid     bool   `dynamodbav:"valid" json:"valid"`
	Validated bool   `dynamodbav:"validated" json:"validated"`
}
