package POST

import (
	"errors"
	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"
	"math/rand"
	"net/http"
	"strconv"

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
		return
	}
	if len(postComment.Commenttext) < 8 {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Yorumunuz 8 karakterden uzun olmalıdır."})
		return
	}

	if res, err := commentCountUser(postComment.Users); res != true {
		g.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
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
func commentCountUser(userID int) (bool, error) {
	var comments = []models.PostComment{}
	db.GetDb().Table("comments").Find(&comments)

	for i, _ := range comments {
		if userid, _ := strconv.Atoi(comments[i].Users); userid == userID {
			return false, errors.New("Bu Makalede Zaten Mevcut Yorumunuz Bulunmakta.")
		}
	}
	return true, nil
}
