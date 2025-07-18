package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nkroshechkin/micro-blog-zero/internal/models"
	"github.com/nkroshechkin/micro-blog-zero/internal/utils"
)

type UserService interface {
	GetAllUser() ([]models.User, error)
	GetUser(id string) (models.User, error)
	CreateUser(name string) (string, error)
}

type userService struct {
	ds *models.DataStructures
}

func NewUserService(ds *models.DataStructures) UserService {
	return &userService{ds: ds}
}

func (s *userService) GetAllUser() ([]models.User, error) {
	users := s.ds.Users
	return users, nil
}

func (s *userService) GetUser(id string) (models.User, error) {
	if id == "" {
		return models.User{}, errors.New("id пустой")
	}

	if user, ok := utils.SearchSliceById(s.ds.Users, id); ok {
		return *user, nil
	}

	return models.User{}, errors.New("пользователь не найден")

}

func (s *userService) CreateUser(name string) (string, error) {

	if name == "" {
		return "", errors.New("имя пользователя пустое")
	}

	likes := []string{}
	newUser := models.User{Id: uuid.New().String(), Username: name, Likes: likes}
	s.ds.Users = append(s.ds.Users, newUser)

	return newUser.Id, nil
}
