package main

import (
	"fmt"
	"net/http"

	"github.com/IraIvanishak/shopping-list-app/config"
	"github.com/IraIvanishak/shopping-list-app/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	// CORS configuration
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type"},
	})

	defer config.DB.Close()

	handler := c.Handler(http.DefaultServeMux)
	router := chi.NewRouter()

	router.Get("/", handlers.GetAllGoodsHandler)
	http.Handle("/", router)

	fmt.Println("starting server on port 8080....")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println("error starting server:", err)
	}
}
