package middleware

import (
	"errors"

	"github.com/umitbasakk/RestApi/models"

	"github.com/umitbasakk/RestApi/db"
)

func FollowValidation(followedbyId int, follower int) (bool, error) {
	var followData = []models.FollowData{}
	db.GetDb().Table("followdata").Find(&followData)

	i := 0
	for i < len(followData) {
		if followData[i].Followerid == follower && followData[i].Followedbyid == followedbyId {
			return false, errors.New("Zaten takip ediyorsun")
		}
		i++
	}

	if followedbyId == follower {
		return false, errors.New("Kendini takip edemezsin")

	}
	return true, nil
}
