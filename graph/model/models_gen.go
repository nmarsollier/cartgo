// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Article struct {
	ID string `json:"id"`
}

func (Article) IsEntity() {}

type Cart struct {
	ID       string         `json:"id"`
	UserID   string         `json:"userId"`
	OrderID  *string        `json:"orderId,omitempty"`
	Articles []*CartArticle `json:"articles"`
	Enabled  bool           `json:"enabled"`
}

func (Cart) IsEntity() {}

type CartArticle struct {
	ArticleID string   `json:"articleId"`
	Article   *Article `json:"article,omitempty"`
	Quantity  int      `json:"quantity"`
	Valid     bool     `json:"valid"`
	Validated bool     `json:"validated"`
}

type Mutation struct {
}

type Query struct {
}
