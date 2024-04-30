package GET

import (
	"net/http"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"

	"github.com/gin-gonic/gin"
)

func (h *GetHandler) GetUser(c *gin.Context) {
	id := c.Param("userid")
	user, resp := h.GetUserTokenParam(id)
	if resp == true {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": resp})
	}
}

func (h *GetHandler) GetUserIdParam(userID int) bool {
	var user = models.User{}
	result := db.GetDb().Find(&user, "userid = ?", userID)
	if result.Row().Err() != nil {
		return false
	}
	return true
}

func (h *GetHandler) GetUserTokenParam(userID string) (models.User, bool) {
	var user = models.User{}
	result := db.GetDb().Find(&user, "token = ?", userID)
	if result.Row().Err() != nil {
		return user, false
	}
	return user, true
}
