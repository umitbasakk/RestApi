package GET

import (
	"ServerRestApi/db"
	"ServerRestApi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *GetHandler) GetAllCategorys(c *gin.Context) {
	var categorys = &[]models.Category{}
	db.GetDb().Table("categorys").Find(categorys)
	c.JSON(http.StatusOK, gin.H{"Categorys": categorys})
}
