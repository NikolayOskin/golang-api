package postController

import (
	"net/http"
    "api/repository/postsRepository"
    "github.com/gorilla/mux"
    "encoding/json"
    "api/models"
    "api/controllers/handler"
)	

func GetAll(w http.ResponseWriter, r *http.Request) {    
    posts := postsRepository.GetAll()
    handler.RespondJSON(w, http.StatusOK, posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    post := postsRepository.GetById(params["id"])
    json.NewEncoder(w).Encode(post)
}

func Create(w http.ResponseWriter, r *http.Request) {
    var post models.Post
    post = models.Post{Slug: r.FormValue("slug"), Title: r.FormValue("title"), Body: r.FormValue("body")}
    result, err := postsRepository.Create(post)
    if err != nil {
        handler.RespondError(w, http.StatusInternalServerError, err.Error())
    }
    handler.RespondJSON(w, http.StatusOK, result)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]
    postsRepository.Delete(id)
}