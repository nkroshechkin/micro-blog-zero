package main

import (
	"fmt"
	"log"

	"github.com/nkroshechkin/micro-blog-zero/internal/handlers"
	"github.com/nkroshechkin/micro-blog-zero/internal/models"
	"github.com/nkroshechkin/micro-blog-zero/pkg/server"
)

func initDS() *models.DataStructures {
	users := make([]models.User, 0)
	posts := make([]models.Post, 0)
	ds := models.DataStructures{Users: users, Posts: posts}
	return &ds
}

func main() {

	ds := initDS()
	routes := handlers.InitRoutes(ds)
	srv := server.Server{}

	fmt.Printf("serverStart port: %s \n", "8080")

	if err := srv.Run("8080", routes); err != nil {
		log.Fatalf("Произошла: %s", err)
	}

}
