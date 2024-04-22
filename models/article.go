package models

import "time"

type Article struct {
	Articleid      string    `json:"Articleid" validation:"req"`
	Imageurl       string    `json:"Imageurl" validation:"req"`
	Title          string    `json:"Title" validation:"req"`
	Createdtime    time.Time `json:"Createdtime" validation:"req"`
	Author         int       `json:"Author" validation:"req"`
	AuthorObject   User
	Category       int    `json:"Category" validation:"req"`
	Articlecontent string `json:"Articlecontent" validation:"req"`
	Comments       []Comment
}
