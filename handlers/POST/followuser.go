package POST

import (
	"net/http"
	"strconv"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/middleware"
	"github.com/umitbasakk/RestApi/models"

	"github.com/gin-gonic/gin"
)

func (h *PostHandler) FollowUser(c *gin.Context) {
	followedby, _ := strconv.Atoi(c.Param("followedby"))
	followed, _ := strconv.Atoi(c.Param("followed")) // takip eden

	if res, err := middleware.FollowValidation(followedby, followed); err != nil || res != true {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	if Follow(followedby, followed) == true {
		c.JSON(http.StatusOK, gin.H{"message": "Başarıyla Takip Edildi."})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Bir sorun oluştu."})

	}

}

func Follow(followedbyId int, follower int) bool {
	var follow = &models.FollowData{}
	follow.Followedbyid = followedbyId
	follow.Followerid = follower
	if res := db.GetDb().Table("followdata").Create(follow); res != nil {
		return false
	}
	return true
}
