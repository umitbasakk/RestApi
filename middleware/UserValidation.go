package middleware

import (
	"github.com/umitbasakk/RestApi/models"

	"github.com/go-playground/validator/v10"
)

func UserValidation(user *models.User) (bool, error) {

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return false, err
	}

	return true, nil
}
