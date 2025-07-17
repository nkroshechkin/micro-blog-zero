package handlers

import (
	"fmt"
	"net/http"

	"github.com/nkroshechkin/micro-blog-zero/internal/models"
	"github.com/nkroshechkin/micro-blog-zero/internal/service"
)

func InitRoutes(ds *models.DataStructures) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "Добро пожаловать на главную страницу!")
	})

	userHandler := initUserHandler(ds)
	router.HandleFunc("GET /users", userHandler.GetUsers)
	router.HandleFunc("POST /register", userHandler.CreateUser)

	postHandler := initPostHandler(ds)
	router.HandleFunc("GET /posts", postHandler.GetPosts)
	router.HandleFunc("POST /posts", postHandler.CreatePosts)
	router.HandleFunc("POST /post/{id}/like", postHandler.LikePosts)

	return router
}

func initUserHandler(ds *models.DataStructures) *UserHandler {
	userService := service.NewUserService(ds)
	userHandler := NewUserHandler(userService)
	return userHandler
}

func initPostHandler(ds *models.DataStructures) *PostHandler {
	postService := service.NewUPostService(ds)
	postHandler := NewPostHandler(postService)
	return postHandler
}
