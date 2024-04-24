package POST

import (
	"ServerRestApi/db"
	"ServerRestApi/middleware"
	"ServerRestApi/models"
	"encoding/json"
	"errors"
	"fmt"

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
		errMarshal, _ := json.Marshal(err.Error())
		g.Writer.Write([]byte(errMarshal))
		return
	} else {
		g.Writer.Write([]byte("Başarıyla Giriş Yapıldı"))
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
