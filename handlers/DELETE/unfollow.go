package DELETE

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"
)

func (h *DeleteHandler) UnFollowUser(g *gin.Context) {
	var unfollowData = &models.FollowUsernameData{}
	g.ShouldBindJSON(&unfollowData)

	if resp := db.GetDb().Table("followdata").Delete(&unfollowData, "Followrequestusername=$1 and targetusername=$2", unfollowData.Followrequestusername, unfollowData.Targetusername); resp.Error != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": resp.Error})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Başarıyla Takipten Çıktınız"})

}
