package middleware

import (
	"errors"

	"github.com/umitbasakk/RestApi/handlers/GET"
	"github.com/umitbasakk/RestApi/models"
)

func ArticleValidation(article *models.ArticlePost) (bool, error) {
	// article validation
	h := &GET.GetHandler{}

	if h.GetUserIdParam(article.Author) != true {
		return false, errors.New("Böyle bir user bulunamadı")
	}
	return true, nil
}
