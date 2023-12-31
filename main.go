package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	Port := "8000"
	if Port != "8000" {
		log.Fatal("Port is not found in the Environmnet")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Route := chi.NewRouter()

	v1Route.HandleFunc("/ready", handlerReadiness)

	router.Mount("/v1", v1Route)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + Port,
	}
	log.Printf("Server Connecting on Port %v", Port)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
		fmt.Println(server)
	}
	fmt.Println(server)
}
