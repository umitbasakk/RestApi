package GET

import (
	"net/http"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"

	"github.com/gin-gonic/gin"
)

func (h *GetHandler) GetAllComments(g *gin.Context) {
	var comments = []models.Comment{}
	db.GetDb().Table("comments").Find(&comments)
	i := 0
	for i < len(comments) {
		comments[i].UserObject = getAuthor(comments[i].Users)
		i++
	}
	g.JSON(http.StatusOK, gin.H{"message": comments})

}
