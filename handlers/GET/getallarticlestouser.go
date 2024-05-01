package GET

import (
	"github.com/gin-gonic/gin"
	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"
	"net/http"
	"strconv"
)

func (h *GetHandler) GetAllArticlesToUser(g *gin.Context) {
	userid := TokentoUserID(g.Param("userid"))
	var currentArticle = []models.Article{}
	db.GetDb().Find(&currentArticle, "author = ?", userid)

	for i, _ := range currentArticle {
		currentArticleID, _ := strconv.Atoi(currentArticle[i].Articleid)
		currentArticle[i].Comments = getComments(currentArticleID)
		currentArticle[i].AuthorObject = getAuthor(currentArticle[i].Author)
	}

	g.JSON(http.StatusOK, currentArticle)
}

func TokentoUserID(token string) int {
	var user = models.User{}
	db.GetDb().Find(&user, "token = ?", token)
	return user.Userid
}
