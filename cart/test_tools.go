package cart

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mock Data
func TestCart() *Cart {
	return &Cart{
		ID:      primitive.NewObjectID(),
		UserId:  "123",
		Enabled: true,
		Created: time.Now(),
		Updated: time.Now(),
		Articles: []*Article{
			{
				ArticleId: "article_1",
				Quantity:  1,
				Valid:     false,
				Validated: false,
			},
			{
				ArticleId: "article_2",
				Quantity:  2,
				Valid:     false,
				Validated: false,
			},
		},
	}
}
