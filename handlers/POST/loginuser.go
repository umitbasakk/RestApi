package POST

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"

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
	res, user, err := HasaUser(LoginUserObj.Username, LoginUserObj.Password)

	if err != nil || res != true {
		g.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	} else {
		tk := &models.Token{Userid: user.Userid, Username: LoginUserObj.Username}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN")))
		user.Token = tokenString
		db.GetDb().Exec("UPDATE users SET token=$1 WHERE userid=$2", tokenString, user.Userid)
		g.JSON(http.StatusOK, gin.H{"message": tokenString})
	}
}

func HasaUser(username string, password string) (bool, models.User, error) {

	var user = models.User{}
	db.GetDb().Table("users").Find(&user, "username = ?", username)

	if user.Username == "" {
		return false, user, errors.New("Böyle bir kullanıcı yok")
	}

	if pass := middleware.ComparePassword(password, user.Password); pass != true {
		return false, user, errors.New("Parola yanlış")
	}
	return true, user, nil

}
