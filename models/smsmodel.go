package models

import "time"

type SMS struct {
	Tophone int    `json:"tophone"`
	Body    string `json:"body"`
}

type SMSRequest struct {
	Userid       int       `json:"Userid"`
	Verifycode   string    `json:"Verifycode"`
	Verifystatus int       `json:"Verifystatus"`
	Updated_at   time.Time `json:"Updated_at"`
}

type SMSPost struct {
	Username   string `json:"Username"`
	Verifycode string `json:"Verifycode"`
}
