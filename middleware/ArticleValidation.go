package middleware

import (
	"errors"

	"github.com/umitbasakk/RestApi/handlers/GET"
	"github.com/umitbasakk/RestApi/models"
)

func ArticleValidation(article *models.ArticlePost) (bool, error) {
	// article validation
	h := &GET.GetHandler{}
	if len(article.Title) < 8 {
		return false, errors.New("Başlık 8 karakterden uzun olmalıdır.")
	}

	if len(article.Articlecontent) < 50 {
		return false, errors.New("Makale 50 karakterden uzun olmalıdır.")
	}

	if h.GetUserIdParam(article.Author) != true {
		return false, errors.New("Böyle bir user bulunamadı")
	}
	return true, nil
}
