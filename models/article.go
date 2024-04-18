package models

import "time"

type Article struct {
	Articleid      string    `json:"Articleid"`
	Imageurl       string    `json:"Imageurl"`
	Title          string    `json:"Title"`
	Createdtime    time.Time `json:"Createdtime"`
	Author         int       `json:"Author"`
	AuthorObject   User
	Category       int    `json:"Category"`
	Articlecontent string `json:"Articlecontent"`
	Comments       []Comment
}
