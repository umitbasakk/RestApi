package GET

import (
	"ServerRestApi/db"
	"ServerRestApi/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *GetHandler) GetFollowersData(c *gin.Context) {
	myid, _ := strconv.Atoi(c.Param("myid"))
	followers, followersby := FollowersData(myid)

	c.String(http.StatusOK, "Takipçilerin:%s \n Takip Ettiğin: %s", followers, followersby)
}

func FollowersData(myid int) (int, int) {
	var followData = []models.FollowData{}
	db.GetDb().Table("followdata").Find(&followData)

	counter := 0
	follower := 0
	followedbyid := 0
	for counter < len(followData) {
		if followData[counter].Followerid == myid {
			follower++
		} else if followData[counter].Followedbyid == myid {
			followedbyid++
		}
		counter++
	}
	return follower, followedbyid
}
