package POST

import (
	"ServerRestApi/db"
	"ServerRestApi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *PostHandler) CreateCategory(c *gin.Context) {

	var currentCategory = &models.Category{}
	c.BindJSON(currentCategory)
	if result := db.GetDb().Table("categorys").Create(currentCategory); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": currentCategory})
}
