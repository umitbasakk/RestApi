package models

import "time"

type User struct {
	Userid             int       `json:"Userid"`
	Profileimageurl    string    `json:"Profileimageurl"`
	Profiledescription string    `json:"Profiledescription"`
	Username           string    `json:"Username"`
	Password           string    `json:"Password"`
	Fullname           string    `json:"Fullname" validate:"required,min=5,max=30"`
	Email              string    `json:"Email"  validate:"email"`
	Mobile             string    `json:"Mobile"`
	Birthday           time.Time `json:"Birthday"`
	Followings         int       `json:"Followings"`
	Followers          int       `json:"Followers"`
	Token              string    `json:"Token"`
}
