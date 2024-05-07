package models

type Bookmark struct {
	Userid    int `json:"Userid"`
	Articleid int `json:"Articleid"`
}

type GetBookmark struct {
	Username  string `json:"Username"`
	Articleid int    `json:"Articleid"`
}
