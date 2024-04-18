package POST

import (
	"ServerRestApi/db"
	"ServerRestApi/middleware"
	"ServerRestApi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *PostHandler) CreateUser(c *gin.Context) {

	var user = &models.User{}
	c.BindJSON(user)
	user.Password, _ = middleware.HashPassword(user.Password)
	if res, err := middleware.UserValidation(user); res != true {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// Auth user
	if result := db.GetDb().Create(user); result.Error != nil {
		c.JSON(200, gin.H{"message": "error"})
	}
	c.JSON(200, gin.H{"message": user})
}
