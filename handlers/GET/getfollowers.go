package GET

import (
	"net/http"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"

	"github.com/gin-gonic/gin"
)

func (h *GetHandler) GetFollowersData(c *gin.Context) {
	username := c.Param("username")
	responseFollowData := models.ResponseFollowData{}
	myfollowers, targetfollowers, followuplist, followersList := FollowersData(username)
	responseFollowData.Followup = targetfollowers
	responseFollowData.Followers = myfollowers
	responseFollowData.Followuplist = followuplist
	responseFollowData.Followerslist = followersList

	c.JSON(http.StatusOK, responseFollowData)
}

func FollowersData(myUsername string) (int, int, []models.ResponseFollowUser, []models.ResponseFollowUser) {
	var followData = []models.FollowUsernameData{}
	db.GetDb().Table("followdata").Find(&followData)

	var followuplist = make([]models.ResponseFollowUser, 0)
	var followersList = make([]models.ResponseFollowUser, 0)

	for _, v := range followData {
		if v.Followrequestusername == myUsername {
			resp, _ := GetUserIDParam(v.Targetusername)
			followuplist = append(followuplist, resp)
		} else if v.Targetusername == myUsername {
			resp, _ := GetUserIDParam(v.Followrequestusername)
			followersList = append(followersList, resp)
		}
	}
	return len(followersList), len(followuplist), followuplist, followersList
}

func GetUserIDParam(userID string) (models.ResponseFollowUser, bool) {
	var tempUser = models.User{}
	var user = models.ResponseFollowUser{}
	result := db.GetDb().Find(&tempUser, "username = ?", userID)
	if result.Row().Err() != nil {
		return user, false
	}
	user = FillValue(tempUser, user)
	return user, true
}

func FillValue(tempUser models.User, user models.ResponseFollowUser) models.ResponseFollowUser {
	user.Username = tempUser.Username
	user.Email = tempUser.Email
	user.Profileimageurl = tempUser.Profileimageurl
	user.Fullname = tempUser.Fullname
	return user
}
