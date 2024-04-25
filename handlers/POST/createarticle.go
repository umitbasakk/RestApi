package POST

import (
	"net/http"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/middleware"
	"github.com/umitbasakk/RestApi/models"

	"github.com/gin-gonic/gin"
)

func (h *PostHandler) CreateArticle(c *gin.Context) {

	var currentArticle = &models.Article{}
	c.BindJSON(currentArticle)
	if result := db.GetDb().Create(currentArticle); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	if result, err := middleware.ArticleValidation(currentArticle); result != true {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": currentArticle})
}
