package GET

import (
	"net/http"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"

	"github.com/gin-gonic/gin"
)

func (h *GetHandler) GetSubCategory(c *gin.Context) {
	var category = &models.TopicCategory{}
	categoryID := c.Param("subcategoryid")
	db.GetDb().Table("topicsubcategorys").Find(category, "categoryid = ?", categoryID)
	c.JSON(http.StatusOK, gin.H{"message": category})
}
