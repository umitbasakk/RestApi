package GET

import (
	"net/http"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"

	"github.com/gin-gonic/gin"
)

func (h *GetHandler) GetUser(c *gin.Context) {
	id := c.Param("userid")
	user, resp := h.GetUserTokenParam(id)
	if resp == true {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": resp})
	}
}

func (h *GetHandler) GetUserIdParam(userID int) bool {
	var user = models.User{}
	result := db.GetDb().Find(&user, "userid = ?", userID)
	if result.Row().Err() != nil {
		return false
	}
	return true
}

func (h *GetHandler) GetUserTokenParam(userID string) (models.ResponseUser, bool) {
	var tempUser = models.User{}
	var user = models.ResponseUser{}
	result := db.GetDb().Find(&tempUser, "token = ?", userID)
	if result.Row().Err() != nil {
		return user, false
	}
	user.Username = tempUser.Username
	user.Birthday = tempUser.Birthday
	user.Email = tempUser.Email
	user.Profileimageurl = tempUser.Profileimageurl
	user.Followers = tempUser.Followers
	user.Followings = tempUser.Followings
	user.Fullname = tempUser.Fullname
	user.Mobile = tempUser.Mobile
	user.Profiledescription = tempUser.Profiledescription
	user.ArticleCount = GetCountArticleOnUser(tempUser.Userid)
	print(user.ArticleCount)
	return user, true
}

func GetCountArticleOnUser(userID int) int {
	var articles = []models.Article{}
	db.GetDb().Find(&articles, "author = ?", userID)
	return len(articles)
}
