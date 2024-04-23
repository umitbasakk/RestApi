package POST

import (
	"ServerRestApi/db"
	"ServerRestApi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *PostHandler) CreateSubCategory(c *gin.Context) {

	var currentSubCategory = &models.TopicCategory{}
	c.BindJSON(currentSubCategory)
	if result := db.GetDb().Table("topicsubcategorys").Create(currentSubCategory); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": currentSubCategory})

}
