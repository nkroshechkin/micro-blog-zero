package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nkroshechkin/micro-blog-zero/internal/handlers"
	"github.com/nkroshechkin/micro-blog-zero/pkg/server"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func main() {
	routes := handlers.InitRoutes()

	srv := server.Server{}

	if err := srv.Run("8080", routes); err != nil {
		log.Fatalf("Произошла: %s", err)
	}

	fmt.Printf("serverStart port: %s \n", "8080")
}
