package cart

import (
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Estuctura basica de del evento
type Cart struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	UserId   string             `bson:"userId"  json:"userId" validate:"required,min=1,max=100"`
	OrderId  string             `bson:"orderId" json:"orderId"`
	Articles []*Article         `bson:"articles"  json:"articles" validate:"required"`
	Enabled  bool               `bson:"enabled" json:"enabled"`
	Created  time.Time          `bson:"created" json:"created"`
	Updated  time.Time          `bson:"updated" json:"updated"`
}

// validateSchema valida la estructura para ser insertada en la db
func (e *Cart) validateSchema() error {
	return validator.New().Struct(e)
}

type Article struct {
	ArticleId string `bson:"articleId" json:"articleId" validate:"required,min=1,max=100"`
	Quantity  int    `bson:"quantity" json:"quantity" validate:"required,min=1,max=100"`
	Valid     bool   `bson:"valid" json:"valid"`
	Validated bool   `bson:"validated" json:"validated"`
}
