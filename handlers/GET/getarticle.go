package GET

import (
	"ServerRestApi/db"
	"ServerRestApi/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *GetHandler) GetArticle(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("articleid"))
	var currentArticle = models.Article{}
	db.GetDb().Find(&currentArticle, "articleid = ?", id)

	currentArticle.Comments = getComments(id)
	currentArticle.AuthorObject = getAuthor(currentArticle.Author)
	c.JSON(http.StatusOK, gin.H{"message": currentArticle})
}

func getComments(articleID int) []models.Comment {
	var comments []models.Comment
	db.GetDb().Find(&comments, "articles = ?", articleID)
	for i, _ := range comments {
		comments[i].UserObject = getAuthor(comments[i].Users)
	}
	return comments
}

func getAuthor(userID int) models.User {
	var user = &models.User{}
	db.GetDb().Find(user, "userid = ?", userID)
	return *user
}
