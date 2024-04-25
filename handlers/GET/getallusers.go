package GET

import (
	"RestApi/db"
	"RestApi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *GetHandler) GetAllUsers(c *gin.Context) {
	var users = &[]models.User{}
	db.GetDb().Find(users)
	c.JSON(http.StatusOK, gin.H{"message": users})
}
