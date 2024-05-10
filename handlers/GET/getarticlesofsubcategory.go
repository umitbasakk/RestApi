package GET

import (
	"github.com/gin-gonic/gin"
	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"
	"net/http"
	"strconv"
)

func (h *GetHandler) GetArticlesofSubcategory(g *gin.Context) {
	id := g.Param("subcategoryid")
	var articles = make([]models.Article, 1)
	if err := db.GetDb().Exec("SELECT * FROM articles INNER JOIN topicsubcategorys_users on articles.author = topicsubcategorys_users.users WHERE topicsubcategorys_users.topicsubcategorys=$1", id).Table("articles").Find(&articles); err.Error != nil {
		g.JSON(http.StatusBadRequest, err.Error)
		return
	}

	for i, _ := range articles {
		currentArticleID, _ := strconv.Atoi(articles[i].Articleid)
		articles[i].Comments = getComments(currentArticleID)
		articles[i].AuthorObject = getAuthor(articles[i].Author)
	}

	g.JSON(http.StatusOK, articles)
}
