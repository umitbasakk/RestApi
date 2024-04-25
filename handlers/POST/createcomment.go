package POST

import (
	"math/rand"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"

	"github.com/gin-gonic/gin"
)

func (h *PostHandler) CreateComment(g *gin.Context) {
	var comment = models.Comment{}
	g.BindJSON(&comment)
	comment.Commentid = rand.Intn(999999999)
	db.GetDb().Create(&comment)

}
