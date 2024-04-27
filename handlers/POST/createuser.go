package POST

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/middleware"
	"github.com/umitbasakk/RestApi/models"

	"github.com/gin-gonic/gin"
)

func (h *PostHandler) CreateUser(c *gin.Context) {

	var user = &models.User{}
	json.NewDecoder(c.Request.Body).Decode(&user)
	user.Password, _ = middleware.HashPassword(user.Password)
	user.Userid = rand.Intn(999999999)
	if res, err := middleware.UserValidation(user); res != true {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// Auth user
	if result := db.GetDb().Create(user); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"message": "Başarıyla Kayıt Olundu"})
}
