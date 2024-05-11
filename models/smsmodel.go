package models

type SMS struct {
	Tophone int    `json:"tophone"`
	Body    string `json:"body"`
}

type SMSRequest struct {
	Userid       int    `json:"Userid"`
	Verifycode   string `json:"Verifycode"`
	Verifystatus int    `json:"Verifystatus"`
}

type SMSPost struct {
	Username   string `json:"Username"`
	Verifycode string `json:"Verifycode"`
}
