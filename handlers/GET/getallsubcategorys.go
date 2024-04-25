package GET

import (
	"net/http"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"

	"github.com/gin-gonic/gin"
)

func (h *GetHandler) GetAllSubCategorys(c *gin.Context) {
	var categorys = &[]models.TopicCategory{}
	db.GetDb().Table("topicsubcategorys").Find(categorys)
	c.JSON(http.StatusOK, gin.H{"message": categorys})
}
