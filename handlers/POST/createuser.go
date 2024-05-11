package POST

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/middleware"
	"github.com/umitbasakk/RestApi/models"

	"github.com/gin-gonic/gin"
)

func (h *PostHandler) CreateUser(c *gin.Context) {

	var user = &models.User{}
	json.NewDecoder(c.Request.Body).Decode(&user)
	if res, err := middleware.UserValidation(user); res != true {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	user.Password, _ = middleware.HashPassword(user.Password)

	// Auth user
	if result := db.GetDb().Create(user); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}

	userVerifyCode := strconv.Itoa(rand.Intn(999999))

	if result := db.GetDb().Table("verifyuser").Exec("INSERT INTO verifyuser(userid,verifycode) VALUES($1,$2)", user.Userid, userVerifyCode); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Başarıyla Kayıt Olundu"})
}
