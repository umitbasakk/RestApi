package POST

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/middleware"
	"github.com/umitbasakk/RestApi/models"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginUserObj struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func (h *PostHandler) LoginUser(g *gin.Context) {
	var LoginUserObj = LoginUserObj{}
	var LoginResponse = models.LoginResponse{}
	g.BindJSON(&LoginUserObj)
	res, user, err := HasaUser(LoginUserObj.Username, LoginUserObj.Password)

	tk := &models.Token{Userid: user.Userid, Username: LoginUserObj.Username}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN")))
	user.Token = tokenString
	db.GetDb().Exec("UPDATE users SET token=$1 WHERE userid=$2", tokenString, user.Userid)

	if err != nil || res != true {
		g.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	} else if resp, code, codeResp := AccountVerify(user); resp != true {

		if codeResp != false {
			_, err := SendSms(code, user.Mobile)

			if err != nil {
				fmt.Printf("error:", err.Error())
			}

		}

		LoginResponse.Token = tokenString
		LoginResponse.Statuscode = 202
		g.JSON(http.StatusAccepted, LoginResponse)
		return
	} else {
		LoginResponse.Token = tokenString
		LoginResponse.Statuscode = 200
		g.JSON(http.StatusOK, LoginResponse)
	}
}
func AccountVerify(user models.User) (bool, string, bool) {
	smsonUser := models.SMSRequest{}
	res := db.GetDb().Table("verifyuser").Find(&smsonUser, "userid = ?", user.Userid)
	userVerifyCode := strconv.Itoa(rand.Intn(999999))

	if res.RowsAffected == 0 {
		db.GetDb().Table("verifyuser").Exec("INSERT INTO verifyuser(userid,verifycode,created_at,updated_at) VALUES($1,$2,$3,$4)", user.Userid, userVerifyCode, time.Now(), time.Now())
		return false, userVerifyCode, true
	}
	if smsonUser.Verifystatus == 0 {
		currentTime := time.Now().Format("2006-01-02-15:04:05")
		smsSendTime := smsonUser.Updated_at.Format("2006-01-02-15:04:05")
		c, _ := time.Parse("2006-01-02-15:04:05", currentTime)
		st, _ := time.Parse("2006-01-02-15:04:05", smsSendTime)
		diffTime := c.Sub(st).Seconds()
		if diffTime < 180 {
			return false, smsonUser.Verifycode, false
		} else {
			db.GetDb().Table("verifyuser").Exec("UPDATE verifyuser SET verifycode=$1 , updated_at=$2 WHERE userid=$3", userVerifyCode, time.Now(), user.Userid)
			return false, userVerifyCode, true
		}
	}

	return true, "", false
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

func SendSms(body string, phone string) (bool, error) {

	client := twilio.NewRestClient()
	toPhone := "+90" + phone
	params := &api.CreateMessageParams{}
	params.SetBody("SMS onay şifreniz:" + body)
	params.SetFrom("+14583022713")
	params.SetTo(toPhone)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return false, err
	} else {
		if resp.Sid != nil {
			return false, err
		} else {
			return true, nil
		}
	}
}
