package GET

import (
	"net/http"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"

	"github.com/gin-gonic/gin"
)

func (h *GetHandler) GetAllUsers(c *gin.Context) {
	var users = &[]models.User{}
	db.GetDb().Find(users)
	c.JSON(http.StatusOK, gin.H{"message": users})
}
