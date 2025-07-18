package models

type User struct {
	Id       string
	Username string
	Likes    []string
}

type ManyUsers struct {
	Users []User `json:"users"`
}
