package models

import "time"

type User struct {
	ProfileURL         string    `json:"profileURL"`
	FullName           string    `json:"fullName"`
	Email              string    `json:"email"`
	Mobile             string    `json:"mobile"`
	Gender             string    `json:"gender"`
	Birthday           time.Time `json:"birthday"`
	Followings         int       `json:"followings"`
	Followers          int       `json:"followers"`
	ProfileDescription string    `json:"profileDescription"`
	WhatsappURL        string    `json:"whatsappURL"`
	FacebookURL        string    `json:"facebookURL"`
	InstagramURL       string    `json:"InstagramURL"`
	TwitterURL         string    `json:"twitterURL"`
}
