package models

import "time"

type User struct {
	ProfileURL         string      `json:"profileURL"`
	Username           string      `json:"username"`
	FullName           string      `json:"fullName"`
	Email              string      `json:"email"`
	Mobile             string      `json:"mobile"`
	Gender             string      `json:"gender"`
	Birthday           time.Time   `json:"birthday"`
	Followings         int         `json:"followings"`
	Followers          int         `json:"followers"`
	ProfileDescription string      `json:"profileDescription"`
	ContactUser        ContactUser `json:"contactUser"`
}
