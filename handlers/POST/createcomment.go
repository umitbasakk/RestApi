package POST

import (
	"ServerRestApi/db"
	"ServerRestApi/models"
	"github.com/gin-gonic/gin"
	"math/rand"
)

func (h *PostHandler) CreateComment(g *gin.Context) {
	var comment = models.Comment{}
	g.BindJSON(&comment)
	comment.Commentid = rand.Intn(999999999)
	db.GetDb().Create(&comment)

}
