package DELETE

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"
	"net/http"
)

func (h *DeleteHandler) DeleteBookmark(g *gin.Context) {
	var bookmark = models.Bookmark{}
	var GetBookmark = models.GetBookmark{}
	g.ShouldBindJSON(&GetBookmark)
	var userID = GetUserFromUsername(GetBookmark.Username)
	bookmark.Userid = userID
	bookmark.Articleid = GetBookmark.Articleid
	if resp, _ := hasBookmark(bookmark); resp != false {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Bir Sorun Oluştu"})
		return
	}
	if resp := db.GetDb().Table("bookmarks").Delete(&bookmark, "userid = ?", bookmark.Userid); resp.Error != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Bir Sorun Oluştu"})
		return
	}
	g.JSON(http.StatusOK, gin.H{"message": "Mark Başarıyla Kaldırıldı"})

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

func hasBookmark(bookmark models.Bookmark) (bool, error) {
	var allbookmarks = []models.Bookmark{}
	if resp := db.GetDb().Table("bookmarks").Find(&allbookmarks, "userid = ?", bookmark.Userid); resp.Error != nil {
		return false, errors.New("Bir sorun oluştu...")
	}
	for _, v := range allbookmarks {
		if v.Userid == bookmark.Userid && v.Articleid == bookmark.Articleid {
			return false, errors.New("Böyle bir mark mevcut değil")
		}
	}
	return true, nil
}
