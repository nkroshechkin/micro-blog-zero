package handlers

import (
	"fmt"
	"net/http"

	"github.com/nkroshechkin/micro-blog-zero/internal/models"
	"github.com/nkroshechkin/micro-blog-zero/internal/service"
)

func InitUserHandler(ds *models.DataStructures) *UserHandler {
	userService := service.NewUserService(ds.Users)
	userHandler := NewUserHandler(userService)
	return userHandler
}

func InitRoutes(ds *models.DataStructures) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "Добро пожаловать на главную страницу!")
	})

	userHandler := InitUserHandler(ds)

	router.HandleFunc("GET /users", userHandler.GetUsers)
	router.HandleFunc("POST /register", userHandler.CreateUser)

	return router
}
