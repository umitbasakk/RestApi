package models

import "time"

type Article struct {
	ImageURL       string    `json:"imageURL"`
	Title          string    `json:"title"`
	CreatedTime    time.Time `json:"createdTime"`
	Author         User      `json:"author"`
	CategoryID     int       `json:"categoryID"`
	ArticleContent string    `json:"articleContent"`
	Comments       []Comment `json:"comments"`
}
