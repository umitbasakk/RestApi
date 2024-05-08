package POST

import (
	"net/http"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/middleware"
	"github.com/umitbasakk/RestApi/models"

	"github.com/gin-gonic/gin"
)

func (h *PostHandler) FollowUser(c *gin.Context) {
	var getFollowData = models.FollowUsernameData{}
	c.ShouldBindJSON(&getFollowData)
	if res, err := middleware.FollowValidation(getFollowData.Followrequestusername, getFollowData.Targetusername); err != nil || res != true {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	if Follow(&getFollowData) == true {
		c.JSON(http.StatusOK, gin.H{"message": "Başarıyla Takip Edildi."})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bir sorun oluştu."})
	}

}

func Follow(follow *models.FollowUsernameData) bool {
	if res := db.GetDb().Table("followdata").Create(&follow); res.Error != nil {
		return false
	}

	if res := db.GetDb().Exec("UPDATE users SET followers=followers+1 WHERE username=$1", follow.Targetusername); res.Error != nil {
		return false
	}
	if res := db.GetDb().Exec("UPDATE users SET followings=followings+1 WHERE username=$1", follow.Followrequestusername); res.Error != nil {
		return false
	}
	return true
}
