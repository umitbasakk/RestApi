package GET

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *GetHandler) GetUser(c *gin.Context) {
	id := c.Param("userid")
	c.JSON(http.StatusOK, gin.H{"message": id})
}
