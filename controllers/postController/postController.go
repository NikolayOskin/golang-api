package postcontroller

import (
	"api/controllers/handler"
	"api/models"
	"api/repository/postsRepository"
	"api/validators"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetAll - get all posts
func GetAll(w http.ResponseWriter, r *http.Request) {
	posts := postsRepository.GetAll()
	handler.RespondJSON(w, http.StatusOK, posts)
}

// GetOne - get exact post
func GetOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	post := postsRepository.GetByIdWithCategory(params["id"])
	// if post == (models.Post{}) {
	// 	handler.RespondError(w, http.StatusNotFound, "404: Not Found")
	// 	return
	// }
	handler.RespondJSON(w, http.StatusOK, post)
}

// Store - store post
func Store(w http.ResponseWriter, r *http.Request) {
	validationError := validators.Handler(w, r)
	if validationError != nil {
		handler.RespondValidationError(w, http.StatusForbidden, validationError)
		return
	}
	var post models.Post
	categoryID, err := strconv.Atoi(r.FormValue("category_id"))
	if err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	post = models.Post{
		Slug:       r.FormValue("slug"),
		Title:      r.FormValue("title"),
		Body:       r.FormValue("body"),
		CategoryID: categoryID,
	}
	result, err := postsRepository.Create(post)
	if err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusOK, result)
}

// Update - update exact post data
func Update(w http.ResponseWriter, r *http.Request) {

}

// Delete - delete exact post
func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if postsRepository.Delete(id) == false {
		handler.RespondError(w, http.StatusNotFound, "404: Not Found")
	}
	handler.RespondJSON(w, http.StatusOK, "Deleted")
}
