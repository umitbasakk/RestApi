package GET

import (
	"ServerRestApi/db"
	"ServerRestApi/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func (h *GetHandler) GetCategory(c *gin.Context) {
	var category = &models.Category{}
	categoryID := c.Param("categoryid")
	db.GetDb().Table("categorys").Find(category, "categoryid = ?", categoryID)
	userJson, _ := json.Marshal(category)
	c.Writer.Write([]byte(userJson))
}
