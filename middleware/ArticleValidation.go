package middleware

import "ServerRestApi/models"

func ArticleValidation(article *models.Article) (bool, error) {
	// article validation
	return true, nil
}
