package categorycontroller

import (
	"api/controllers/handler"
	"api/models"
	"api/repository/categoryrepository"
	"net/http"

	"github.com/gorilla/mux"
)

// GetOne - get category with related posts
func GetOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	category := categoryrepository.GetByIDWithPosts(params["id"])
	handler.RespondJSON(w, http.StatusOK, category)
}

// Store - store category in db
func Store(w http.ResponseWriter, r *http.Request) {
	category := &models.Category{Slug: r.FormValue("slug"), Title: r.FormValue("title")}
	result, err := categoryrepository.Store(category)
	if err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusOK, result)
}

// Update - update existed category
func Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	category := &models.Category{Slug: r.FormValue("slug"), Title: r.FormValue("title")}
	result, err := categoryrepository.UpdateByID(params["id"], category)
	if err != nil {
		handler.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondJSON(w, http.StatusOK, result)
}

func Delete() {

}
