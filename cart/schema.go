package cart

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Estuctura basica de del evento
type Cart struct {
	ID       string     `json:"id"`
	UserId   string     `json:"userId" validate:"required,min=1,max=100"`
	OrderId  string     `json:"orderId"`
	Articles []*Article `json:"articles" validate:"required"`
	Enabled  bool       `json:"enabled"`
	Created  time.Time  `json:"created"`
	Updated  time.Time  `json:"updated"`
}

// validateSchema valida la estructura para ser insertada en la db
func (e *Cart) validateSchema() error {
	return validator.New().Struct(e)
}

type Article struct {
	ArticleId string `json:"articleId" validate:"required,min=1,max=100"`
	Quantity  int    `json:"quantity" validate:"required,min=1,max=100"`
	Valid     bool   `json:"valid"`
	Validated bool   `json:"validated"`
}
