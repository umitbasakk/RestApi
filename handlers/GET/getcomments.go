package GET

import (
	"net/http"
	"strconv"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"

	"github.com/gin-gonic/gin"
)

func (h *GetHandler) GetComments(g *gin.Context) {
	articleID, _ := strconv.Atoi(g.Param("articleid"))
	var comments = []models.Comment{}
	db.GetDb().Table("comments").Find(&comments, "articles = ?", articleID)
	i := 0
	for i < len(comments) {
		comments[i].UserObject = getAuthor(comments[i].Users)
		i++
	}
	g.JSON(http.StatusOK, gin.H{"message": comments})

}
