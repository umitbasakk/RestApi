package GET

import (
	"RestApi/db"
	"RestApi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *GetHandler) GetAllCategorys(c *gin.Context) {
	var categorys = &[]models.Category{}
	db.GetDb().Table("categorys").Find(categorys)
	c.JSON(http.StatusOK, gin.H{"Categorys": categorys})
}
