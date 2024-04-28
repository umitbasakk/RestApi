package POST

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/middleware"
	"github.com/umitbasakk/RestApi/models"

	"github.com/gin-gonic/gin"
)

type LoginUserObj struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func (h *PostHandler) LoginUser(g *gin.Context) {
	var LoginUserObj = LoginUserObj{}
	g.BindJSON(&LoginUserObj)
	fmt.Printf("Username: %s Password: %s", LoginUserObj.Username, LoginUserObj.Password)
	res, err := HasaUser(LoginUserObj.Username, LoginUserObj.Password)

	if err != nil || res != true {
		g.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	} else {
		g.JSON(http.StatusOK, gin.H{"message": "Başarıyla Giriş Yapıldı."})
	}
}

func HasaUser(username string, password string) (bool, error) {

	var user = models.User{}
	db.GetDb().Table("users").Find(&user, "username = ?", username)

	if user.Username == "" {
		return false, errors.New("Böyle bir kullanıcı yok")
	}

	if pass := middleware.ComparePassword(password, user.Password); pass != true {
		return false, errors.New("Parola yanlış")
	}
	return true, nil

}
