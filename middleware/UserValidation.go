package middleware

import (
	"errors"

	"github.com/umitbasakk/RestApi/models"
)

func UserValidation(user *models.User) (bool, error) {

	if len(user.Fullname) < 6 || user.Fullname == "" {
		return false, errors.New("Adınızı boş veya 6 Karakterden az olmamalıdır")
	}

	if len(user.Username) < 6 || user.Username == "" {
		return false, errors.New("Kullanıcı Adınızı boş veya 6 Karakterden az olmamalıdır")
	}

	if len(user.Password) < 8 {
		return false, errors.New("Parolanız 8 karakterden fazla olmalıdır.")
	}

	return true, nil
}
