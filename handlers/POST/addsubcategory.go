package POST

import (
	"github.com/gin-gonic/gin"
	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"
	"net/http"
	"strconv"
)

func (h *PostHandler) AddSubCategoryOnUser(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Param("userid"))
	var topicList = []models.TopicCategory{}
	c.ShouldBindJSON(&topicList)
	var matchList = make([]models.MatchTopic, len(topicList))
	for i := 0; i < len(topicList); i++ {
		val, _ := strconv.Atoi(topicList[i].Categoryid)
		matchList[i] = models.MatchTopic{Topicsubcategorys: val, Users: userid}
	}
	for i := 0; i < len(matchList); i++ {
		db.GetDb().Table("topicsubcategorys_users").Create(matchList[i])
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})

}
