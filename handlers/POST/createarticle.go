package POST

import (
	"encoding/json"
	"github.com/umitbasakk/RestApi/middleware"
	"net/http"
	"time"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"

	"github.com/gin-gonic/gin"
)

func (h *PostHandler) CreateArticle(c *gin.Context) {
	var token = c.Param("userid")
	var currentArticle = models.ArticlePost{}
	json.NewDecoder(c.Request.Body).Decode(&currentArticle)
	currentArticle.Createdtime = time.Now()

	if result, err := middleware.ArticleValidation(&currentArticle); result != true {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user = models.User{}
	db.GetDb().Find(&user, "token = ?", token)
	currentArticle.Author = user.Userid

	if result := db.GetDb().Table("articles").Create(currentArticle); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Makale Başarıyla Paylaşıldı..."})
}
