package POST

import (
	"github.com/umitbasakk/RestApi/handlers/GET"
	"github.com/umitbasakk/RestApi/middleware"
	"net/http"
	"time"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"

	"github.com/gin-gonic/gin"
)

func (h *PostHandler) CreateArticle(c *gin.Context) {
	var token = c.Param("userid")
	var userid = GET.TokentoUserID(token)
	var currentArticle = &models.ArticlePost{}
	c.ShouldBindJSON(currentArticle)
	currentArticle.Createdtime = time.Now()
	currentArticle.Author = userid
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
