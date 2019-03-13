package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "api/controllers/postController"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/posts", postController.GetAll).Methods("GET")
    router.HandleFunc("/post/{id}", postController.GetPost).Methods("GET")
    router.HandleFunc("/posts/create", postController.Create).Methods("POST")
    router.HandleFunc("/posts/{id}", postController.Delete).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8080", router))
}