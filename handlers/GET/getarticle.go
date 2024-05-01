package GET

import (
	"net/http"
	"strconv"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"

	"github.com/gin-gonic/gin"
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

func getAuthor(userID int) models.ResponseUser {
	var tempUser = models.User{}
	var user = models.ResponseUser{}
	db.GetDb().Find(&tempUser, "userid = ?", userID)
	FillUser(&user, &tempUser)
	return user
}

func FillUser(user *models.ResponseUser, tempUser *models.User) {
	user.Username = tempUser.Username
	user.Birthday = tempUser.Birthday
	user.Email = tempUser.Email
	user.Profileimageurl = tempUser.Profileimageurl
	user.Followers = tempUser.Followers
	user.Followings = tempUser.Followings
	user.Fullname = tempUser.Fullname
	user.Gender = tempUser.Gender
	user.Mobile = tempUser.Mobile
	user.Profiledescription = tempUser.Profiledescription
	user.ArticleCount = GetCountArticleOnUser(tempUser.Userid)
}
