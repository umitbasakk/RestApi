package models

type Comment struct {
	Commentid   int    `json:"Commentid"`
	UserObject  User   `json:"UserObject"`
	Users       int    `json:"Users"`
	Commenttext string `json:"Commenttext"`
	Articles    string `json:"Articles"`
}
