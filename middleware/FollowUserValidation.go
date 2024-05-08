package middleware

import (
	"errors"

	"github.com/umitbasakk/RestApi/models"

	"github.com/umitbasakk/RestApi/db"
)

func FollowValidation(Followrequestusername string, Targetusername string) (bool, error) {
	var followData = []models.FollowUsernameData{}
	db.GetDb().Table("followdata").Find(&followData)

	i := 0
	for i < len(followData) {
		if followData[i].Targetusername == Targetusername && followData[i].Followrequestusername == Followrequestusername {
			return false, errors.New("Zaten takip ediyorsun")
		}
		i++
	}

	if Followrequestusername == Targetusername {
		return false, errors.New("Kendini takip edemezsinx")

	}
	return true, nil
}
