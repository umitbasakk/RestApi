package GET

import (
	"net/http"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"

	"github.com/gin-gonic/gin"
)

func (h *GetHandler) GetUser(c *gin.Context) {
	id := c.Param("userid")
	c.JSON(http.StatusOK, gin.H{"message": id})
}

func (h *GetHandler) GetUserIdParam(userID int) bool {
	result := db.GetDb().Find(&models.User{}, "userid = ?", userID)
	if result.Row().Err() != nil {
		return false
	}
	return true
}
