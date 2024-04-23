package middleware

import (
	"ServerRestApi/db"
	"ServerRestApi/models"
	"errors"
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
