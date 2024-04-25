package POST

import (
	"RestApi/db"
	"RestApi/models"
	"math/rand"

	"github.com/gin-gonic/gin"
)

func (h *PostHandler) CreateComment(g *gin.Context) {
	var comment = models.Comment{}
	g.BindJSON(&comment)
	comment.Commentid = rand.Intn(999999999)
	db.GetDb().Create(&comment)

}
