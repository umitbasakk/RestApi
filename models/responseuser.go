package models

import "time"

type ResponseUser struct {
	Username           string    `json:"Username"`
	Profileimageurl    string    `json:"Profileimageurl"`
	ArticleCount       int       `json:"ArticleCount"`
	Fullname           string    `json:"Fullname" validate:"required,min=5,max=30"`
	Email              string    `json:"Email"  validate:"email"`
	Mobile             string    `json:"Mobile"`
	Gender             string    `json:"Gender"`
	Birthday           time.Time `json:"Birthday"`
	Followings         int       `json:"Followings"`
	Followers          int       `json:"Followers"`
	Profiledescription string    `json:"Profiledescription"`
}

type ResponseFollowUser struct {
	Username        string `json:"Username"`
	Profileimageurl string `json:"Profileimageurl"`
	Fullname        string `json:"Fullname" validate:"required,min=5,max=30"`
	Email           string `json:"Email"  validate:"email"`
}

type ResponseFollowData struct {
	Followup      int                  `json:"Followup"`
	Followers     int                  `json:"Followers"`
	Followuplist  []ResponseFollowUser `json:"Followuplist"`
	Followerslist []ResponseFollowUser `json:"Followerslist"`
}
