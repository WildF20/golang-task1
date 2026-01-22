package handler

import (
	"net/http"
	"encoding/json"

	"golang-task1/model"
)

var categories = []model.Category{}

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	return
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	return
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	return
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	return
}