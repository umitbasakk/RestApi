package models

type Comment struct {
	User       User   `json:"user"`
	CommenText string `json:"commenText"`
}
