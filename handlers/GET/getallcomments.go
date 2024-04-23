package GET

import (
	"ServerRestApi/db"
	"ServerRestApi/models"
	"github.com/gin-gonic/gin"
	"net/http"
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
