package POST

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"
)

func (h *PostHandler) CreateBookmark(g *gin.Context) {
	var bookmark = models.Bookmark{}
	var GetBookmark = models.GetBookmark{}
	g.ShouldBindJSON(&GetBookmark)
	var userID = GetUserFromUsername(GetBookmark.Username)
	bookmark.Userid = userID
	bookmark.Articleid = GetBookmark.Articleid
	if resp, err := hasBookmark(bookmark); resp != true {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if res := db.GetDb().Table("bookmarks").Create(&bookmark); res.Error != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Bir Sorun Oluştu..."})
		return
	}
	g.JSON(http.StatusOK, gin.H{"message": "Başarıyla Eklendi"})
}

func hasBookmark(bookmark models.Bookmark) (bool, error) {
	var allbookmarks = []models.Bookmark{}
	db.GetDb().Table("bookmarks").Find(&allbookmarks, "userid = ?", bookmark.Userid)

	for _, v := range allbookmarks {
		if v.Userid == bookmark.Userid && v.Articleid == bookmark.Articleid {
			return false, errors.New("Böyle bir mark zaten mevcut")
		}
	}
	return true, nil
}
