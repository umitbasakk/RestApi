package middleware

import (
	"errors"

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

	for _, v := range *users {
		if v.Email == user.Email {
			return false, errors.New("Bu Kullanıcı adı başka bir kullanıcı tarafından kullanılıyor")
		}
		if v.Mobile == user.Mobile {
			return false, errors.New("Bu numara başkası tarafından kullanılmaktadır.")
		}
	}

	return true, nil
}
