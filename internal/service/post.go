package service

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/nkroshechkin/micro-blog-zero/internal/models"
	"github.com/nkroshechkin/micro-blog-zero/internal/utils"
)

type PostService interface {
	GetAllPost() ([]models.Post, error)
	GetPost(id string) (models.Post, error)
	CreatePost(authorId string, text string) (string, error)
	LikePost(userId string, postId string) (string, error)
}

type postService struct {
	ds *models.DataStructures
}

func NewUPostService(ds *models.DataStructures) PostService {
	return &postService{ds: ds}
}

func (p *postService) GetAllPost() ([]models.Post, error) {
	return p.ds.Posts, nil
}

func (p *postService) GetPost(id string) (models.Post, error) {
	if id == "" {
		return models.Post{}, errors.New("id пустой")
	}
	if post, ok := utils.SearchSliceById(p.ds.Posts, id); ok {
		return *post, nil
	}
	return models.Post{}, errors.New("пост не найден")
}

func (p *postService) CreatePost(authorId string, text string) (string, error) {

	if _, ok := utils.SearchSliceById(p.ds.Users, authorId); !ok {
		return "", errors.New("некоректный пользователь")
	}

	likes := []string{}
	newPost := models.Post{Id: uuid.New().String(), AuthorId: authorId, Text: text, LikeList: likes}
	p.ds.Posts = append(p.ds.Posts, newPost)

	return newPost.Id, nil
}

func (p *postService) LikePost(userId string, postId string) (string, error) {

	user, userFound := utils.SearchSliceById(p.ds.Users, userId)
	post, postFound := utils.SearchSliceById(p.ds.Posts, postId)
	if !userFound {
		return "", errors.New("некоректный пользователь")
	}

	if !postFound {
		return "", errors.New("некоректный пост")
	}

	for _, item := range post.LikeList {
		if item == userId {
			user.Likes = utils.SliceFilter(user.Likes, func(item string) bool {
				return item != userId
			})
			post.LikeList = utils.SliceFilter(post.LikeList, func(item string) bool {
				return item != userId
			})

			return "ok", nil
		}
	}

	user.Likes = append(user.Likes, postId)
	post.LikeList = append(post.LikeList, userId)

	fmt.Println(post.LikeList, p.ds.Posts)

	return "ok", nil

}
