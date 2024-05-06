package POST

import (
	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *PostHandler) CreateComment(g *gin.Context) {
	var getComment = models.PostComment{}
	g.BindJSON(&getComment)
	var postComment = models.Comment{}

	username := getComment.Users

	postComment.Commentid = rand.Intn(999999999)
	postComment.Users = GetUserFromUsername(username)
	postComment.Articles = getComment.Articles
	postComment.Commenttext = getComment.Commenttext

	if postComment.Users == -1 {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Böyle bir kullanıcı yok"})
	}

	db.GetDb().Create(&postComment)
	g.JSON(http.StatusOK, gin.H{"message": "Yorum Başarıyla Gönderildi."})
}

func GetUserFromUsername(username string) int {
	var users = []models.User{}
	db.GetDb().Find(&users)
	for i, _ := range users {
		if users[i].Username == username {
			return users[i].Userid
		}
	}
	return -1
}
