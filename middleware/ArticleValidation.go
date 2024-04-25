package middleware

import (
	"RestApi/handlers/GET"
	"RestApi/models"
	"errors"

	validator2 "github.com/go-playground/validator/v10"
)

func ArticleValidation(article *models.Article) (bool, error) {
	// article validation

	validator := validator2.New()
	h := &GET.GetHandler{}
	if err := validator.Struct(article); err != nil {
		return false, err
	}

	if h.GetUserIdParam(article.Author) != true {
		return false, errors.New("Böyle bir user bulunamadı")
	}
	return true, nil
}
