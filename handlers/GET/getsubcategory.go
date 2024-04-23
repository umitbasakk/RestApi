package GET

import (
	"ServerRestApi/db"
	"ServerRestApi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *GetHandler) GetSubCategory(c *gin.Context) {
	var category = &models.TopicCategory{}
	categoryID := c.Param("subcategoryid")
	db.GetDb().Table("topicsubcategorys").Find(category, "categoryid = ?", categoryID)
	c.JSON(http.StatusOK, gin.H{"message": category})
}
