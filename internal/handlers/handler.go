package handlers

import (
	"fmt"
	"net/http"
)

func InitRoutes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "Добро пожаловать на главную страницу!")
	})

	return router
}
