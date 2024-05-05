package POST

import (
	"github.com/umitbasakk/RestApi/middleware"
	"net/http"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/handlers/GET"
	"github.com/umitbasakk/RestApi/models"

	"github.com/gin-gonic/gin"
)

func (h *PostHandler) CreateArticle(c *gin.Context) {

	var userId = GET.TokentoUserID(c.Param("userid"))
	var currentArticle = &models.ArticlePost{}
	c.BindJSON(currentArticle)
	currentArticle.Author = userId

	if result, err := middleware.ArticleValidation(currentArticle); result != true {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.GetDb().Table("articles").Create(currentArticle); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Makale Başarıyla Paylaşıldı..."})
}
