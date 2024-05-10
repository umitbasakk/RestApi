package GET

import (
	"github.com/gin-gonic/gin"
	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"
	"net/http"
	"strconv"
)

type TopicUserArticle struct {
	Topicsubcategorys int `json:"Topicsubcategorys"`
	Users             int `json:"Users"`
}

func (h *GetHandler) GetArticlesofSubcategory(g *gin.Context) {
	id := g.Param("subcategoryid")
	var topicList = make([]TopicUserArticle, 0)
	var articleList = make([]models.Article, 0)
	var articresponseListle = make([]models.Article, 0)

	if err := db.GetDb().Table("topicsubcategorys_users").Find(&topicList, "topicsubcategorys = ? ", id); err.Error != nil {
		g.JSON(http.StatusBadRequest, err.Error)
		return
	}

	if err := db.GetDb().Table("articles").Find(&articleList); err.Error != nil {
		g.JSON(http.StatusBadRequest, err.Error)
		return
	}
	for _, v := range topicList {
		for j, z := range articleList {
			if v.Users == z.Author {
				articresponseListle = append(articresponseListle, articleList[j])
			}
		}
	}

	for i, _ := range articresponseListle {
		currentArticleID, _ := strconv.Atoi(articresponseListle[i].Articleid)
		articresponseListle[i].Comments = getComments(currentArticleID)
		articresponseListle[i].AuthorObject = getAuthor(articresponseListle[i].Author)
	}

	g.JSON(http.StatusOK, articresponseListle)
}
