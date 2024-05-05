package middleware

import (
	"errors"
	"strconv"
	"strings"

	"github.com/umitbasakk/RestApi/db"
	"github.com/umitbasakk/RestApi/models"
)

func UserValidation(user *models.User) (bool, error) {

	var users = &[]models.User{}

	db.GetDb().Find(users)
	if len(user.Fullname) < 6 || user.Fullname == "" {
		return false, errors.New("Adınızı boş veya 6 Karakterden az olmamalıdır")
	}

	if len(user.Username) < 6 || user.Username == "" {
		return false, errors.New("Kullanıcı Adınızı boş veya 6 Karakterden az olmamalıdır")
	}

	if len(user.Password) < 8 {
		return false, errors.New("Parolanız 8 karakterden fazla olmalıdır.")
	}
	if len(user.Profiledescription) < 8 {
		return false, errors.New("Profil Açıklamanız 50 karakterden fazla olmalıdır.")
	}
	if len(user.Mobile) != 10 {
		return false, errors.New("Numaranın başında 0 olmayacak şekilde yazınız.")
	}
	if _, err := strconv.ParseInt(user.Mobile, 10, 64); err != nil {
		return false, errors.New("Numaranız sadece rakamlardan oluşmalıdır")
	}

	for _, v := range *users {
		if v.Email == user.Email {
			return false, errors.New("Bu Email başka bir kullanıcı tarafından kullanılıyor")
		}
		if v.Username == user.Username {
			return false, errors.New("Bu Kullanıcı adı başka bir kullanıcı tarafından kullanılıyor")
		}
		if v.Mobile == user.Mobile {
			return false, errors.New("Bu numara başkası tarafından kullanılmaktadır.")
		}
	}

	if !strings.Contains(user.Email, "@gmail.com") {
		return false, errors.New("Lütfen geçerli bir gmail adresi giriniz.")
	}

	return true, nil
}
