package models

type ContactUser struct {
	WhatsappURL  string `json:"whatsappURL"`
	FacebookURL  string `json:"facebookURL"`
	InstagramURL string `json:"InstagramURL"`
	TwitterURL   string `json:"twitterURL"`
	Address      string `json:"address"`
	Country      string `json:"countryb"`
}
