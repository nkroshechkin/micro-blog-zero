package main

import (
	"fmt"
	"log"

	"github.com/nkroshechkin/micro-blog-zero/internal/handlers"
	"github.com/nkroshechkin/micro-blog-zero/internal/models"
	"github.com/nkroshechkin/micro-blog-zero/pkg/server"
)

func initDS() *models.DataStructures {
	users := []models.User{}
	ds := models.DataStructures{Users: &users}
	return &ds
}

func main() {

	ds := initDS()
	routes := handlers.InitRoutes(ds)
	srv := server.Server{}

	if err := srv.Run("8080", routes); err != nil {
		log.Fatalf("Произошла: %s", err)
	}

	fmt.Printf("serverStart port: %s \n", "8080")
}
