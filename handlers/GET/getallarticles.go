package GET

import (
	"RestApi/db"
	"RestApi/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *GetHandler) GetAllArticles(c *gin.Context) {
	var currentArticle = []models.Article{}
	db.GetDb().Find(&currentArticle)

	for i, _ := range currentArticle {
		currentArticleID, _ := strconv.Atoi(currentArticle[i].Articleid)
		currentArticle[i].Comments = getComments(currentArticleID)
		currentArticle[i].AuthorObject = getAuthor(currentArticle[i].Author)
	}

	c.JSON(http.StatusOK, gin.H{"message": currentArticle})
}
