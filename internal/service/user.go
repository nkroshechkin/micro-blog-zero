package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nkroshechkin/micro-blog-zero/internal/models"
)

type UserService interface {
	GetAllUser() (*[]models.User, error)
	GetUser(id string) (models.User, error)
	CreateUser(name string) (string, error)
}

type userService struct {
	users *[]models.User
}

func NewUserService(users *[]models.User) UserService {
	return &userService{users: users}
}

func (s *userService) GetAllUser() (*[]models.User, error) {
	users := s.users
	return users, nil
}

func (s *userService) GetUser(id string) (models.User, error) {
	if id == "" {
		return models.User{}, errors.New("id пустой")
	}

	for _, item := range *s.users {
		if itemID := item.Id; itemID == id {
			return item, nil
		}
	}

	return models.User{}, errors.New("пользователь не найден")

}

func (s *userService) CreateUser(name string) (string, error) {
	likes := []string{}

	newUser := models.User{Id: uuid.New().String(), Username: name, Likes: likes}

	*s.users = append(*s.users, newUser)

	return newUser.Id, nil
}
