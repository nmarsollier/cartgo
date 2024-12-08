package cart

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Mock Data
func TestCart() *Cart {
	return &Cart{
		ID:      uuid.NewV4().String(),
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
