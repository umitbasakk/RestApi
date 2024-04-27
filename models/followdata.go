package models

type FollowData struct {
	Followerid   int `json:"Followerid"`
	Followedbyid int `json:"Followedbyid"`
}
