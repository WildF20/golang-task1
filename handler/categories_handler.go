package handler

import (
	"net/http"
	"encoding/json"

	"golang-task1/model"
	"golang-task1/model/structs"
)

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories := []model.Category{}

	response := structs.SuccessResponse{
		Status: true,
		Message: "Success",
		Data: categories,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		errResponse := structs.ErrorResponse{
			Status: false,
			Message: err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errResponse)
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