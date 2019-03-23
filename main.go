package main

import (
	"api/controllers/categorycontroller"
	"api/controllers/postcontroller"
	"api/db"
	"api/db/migrations"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	if _, err := db.Open(); err != nil {
		panic(err.Error())
	}
	defer db.Close()

	migrations.Migrate()

	router := mux.NewRouter()

	// posts
	router.HandleFunc("/posts", postcontroller.GetAll).Methods("GET")
	router.HandleFunc("/posts/{id}", postcontroller.GetOne).Methods("GET")
	router.HandleFunc("/posts", postcontroller.Store).Methods("POST")
	router.HandleFunc("/posts/{id}", postcontroller.Update).Methods("PUT")
	router.HandleFunc("/posts/{id}", postcontroller.Delete).Methods("DELETE")

	// categories
	router.HandleFunc("/categories/{id}", categorycontroller.GetOne).Methods("GET")
	router.HandleFunc("/categories", categorycontroller.Store).Methods("POST")
	router.HandleFunc("/categories/{id}", categorycontroller.Update).Methods("PUT")
	//router.HandleFunc("/categories/{id}", categorycontroller.Delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8084", router))
}
