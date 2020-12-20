package main

import (
	"mindteck/upload"
	"mindteck/retrival"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/rs/cors"
)

func main(){
	router := mux.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{
			"Origin", "Authorization", "Access-Control-Allow-Origin",
			"Access-Control-Allow-Header", "Accept",
			"Content-Type", "X-CSRF-Token",
		},
		ExposedHeaders: []string{
			"Content-Length", "Access-Control-Allow-Origin", "Origin",
		},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router.Use(cors.Handler)
	router.HandleFunc("/upload", upload.UploadImage)
	router.HandleFunc("/retrieve", view.RetrieveImages)
	router.HandleFunc("/search", view.SearchByID)
	http.ListenAndServe(":9195", router)
}