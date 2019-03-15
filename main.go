package main

import (
	"api/controllers/postController"
	"api/db"
	"api/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	if _, err := db.Open(); err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.DB.DropTableIfExists("posts")
	db.DB.AutoMigrate(&models.Post{})

	router := mux.NewRouter()
	router.HandleFunc("/posts", postController.GetAll).Methods("GET")
	router.HandleFunc("/posts/{id}", postController.GetPost).Methods("GET")
	router.HandleFunc("/posts/create", postController.Create).Methods("POST")
	router.HandleFunc("/posts/{id}", postController.Update).Methods("PUT")
	router.HandleFunc("/posts/{id}", postController.Delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
