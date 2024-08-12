package tests

import (
	"time"

	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/security"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mock Data
func TestCart() *cart.Cart {
	return &cart.Cart{
		ID:      primitive.NewObjectID(),
		UserId:  "123",
		Enabled: true,
		Created: time.Now(),
		Updated: time.Now(),
		Articles: []*cart.Article{
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

func TestUser() *security.User {
	return &security.User{
		ID:          primitive.NewObjectID().String(),
		Login:       "Login",
		Name:        "Name",
		Permissions: []string{"user", "other"},
	}
}
