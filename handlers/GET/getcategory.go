package GET

import (
	"ServerRestApi/db"
	"ServerRestApi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *GetHandler) GetCategory(c *gin.Context) {
	var category = &models.Category{}
	categoryID := c.Param("categoryid")
	db.GetDb().Table("categorys").Find(category, "categoryid = ?", categoryID)
	c.JSON(http.StatusOK, gin.H{"message": category})
}
