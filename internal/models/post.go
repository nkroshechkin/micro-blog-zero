package models

type Post struct {
	Id       string
	AuthorId string
	Text     string
	LikeList *[]string
}
