package models

import (
	"time"
)

type Article struct {
	Articleid      string    `json:"Articleid" validation:"req"`
	Imageurl       string    `json:"Imageurl" validation:"req"`
	Title          string    `json:"Title" validation:"req"`
	Createdtime    time.Time `json:"Createdtime"`
	Author         int       `json:"Author" validation:"req"`
	AuthorObject   ResponseUser
	Category       int       `json:"Category" validation:"req"`
	Articlecontent string    `json:"Articlecontent" validation:"req"`
	Comments       []Comment `json:"Comments"`
	Allowcomments  int       `json:"Allowcomments"`
}

type ArticlePost struct {
	Articleid      string    `json:"Articleid"`
	Imageurl       string    `json:"Imageurl"`
	Title          string    `json:"Title"`
	Createdtime    time.Time `json:"Createdtime"`
	Author         int       `json:"Author"`
	Category       int       `json:"Category"`
	Articlecontent string    `json:"Articlecontent"`
	Allowcomments  int       `json:"Allowcomments"`
}
