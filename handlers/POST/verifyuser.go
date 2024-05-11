package POST

import (
	"github.com/gin-gonic/gin"
	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"
	"net/http"
)

func (h *PostHandler) VerifyUser(g *gin.Context) {
	postSmsVerify := models.SMSPost{}
	SMSVerify := models.SMSRequest{}
	g.ShouldBindJSON(&postSmsVerify)
	userID := GetUserFromUsername(postSmsVerify.Username)
	db.GetDb().Table("verifyuser").Find(&SMSVerify, "userid = ?", userID)

	if postSmsVerify.Verifycode == SMSVerify.Verifycode {
		db.GetDb().Table("verifyuser").Exec("UPDATE verifyuser SET verifystatus=$1 WHERE userid=$2", 1, userID)
		g.JSON(http.StatusBadRequest, gin.H{"message": "Doğrulama İşleminiz Başarılı.Yönlendiriliyorsunuz..."})
		return

	} else {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Doğrulama Kodunuz Hatalı"})
	}

}
