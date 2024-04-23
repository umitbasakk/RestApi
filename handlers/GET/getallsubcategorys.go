package GET

import (
	"ServerRestApi/db"
	"ServerRestApi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *GetHandler) GetAllSubCategorys(c *gin.Context) {
	var categorys = &[]models.TopicCategory{}
	db.GetDb().Table("topicsubcategorys").Find(categorys)
	c.JSON(http.StatusOK, gin.H{"message": categorys})
}
