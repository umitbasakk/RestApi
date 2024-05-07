package POST

import (
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
	if res := db.GetDb().Table("bookmarks").Create(&bookmark); res.Error != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Bir Sorun Oluştu..."})
		return
	}
	g.JSON(http.StatusOK, gin.H{"message": "Başarıyla Eklendi"})
}
