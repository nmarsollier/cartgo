package cart

type ValidationEvent struct {
	ReferenceId string `json:"referenceId" example:"UserId"`
	ArticleId   string `json:"articleId" example:"ArticleId"`
	Valid       bool   `json:"valid" example:"true"`
}

func ProcessArticleData(data *ValidationEvent, deps ...interface{}) error {
	cart, err := findByUserId(data.ReferenceId, deps...)
	if err != nil {
		return err
	}

	for _, a := range cart.Articles {
		if a.ArticleId == data.ArticleId {
			a.Validated = true
			a.Valid = data.Valid
		}
	}

	err = replace(cart, deps...)
	if err != nil {
		return err
	}

	return nil
}
