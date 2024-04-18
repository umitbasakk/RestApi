package GET

import (
	"ServerRestApi/db"
	"ServerRestApi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *GetHandler) GetAllUsers(c *gin.Context) {
	var users = &[]models.User{}
	db.GetDb().Find(users)
	c.JSON(http.StatusOK, gin.H{"message": users})
}
