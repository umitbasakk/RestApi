package GET

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"
)

func (h *GetHandler) GetBookmark(g *gin.Context) {
	username := g.Param("username")
	userID := GetUserFromUsername(username)
	var bookmarks = []models.Bookmark{}
	if userID == -1 {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Böyle Bir Kullanıcı Yok"})
		return
	}
	if res := db.GetDb().Table("bookmarks").Find(&bookmarks, "userid = ?", userID); res.Error != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": bookmarks})

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
