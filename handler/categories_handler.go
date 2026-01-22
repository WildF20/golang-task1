package handler

import (
	"strings"
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/go-playground/validator/v10"

	"golang-task1/model"
	"golang-task1/model/structs"
)

var (
	categories = []model.Category{}
)

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
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
	var newCategory model.Category

	if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
		errResponse := structs.ErrorResponse{
			Status: false,
			Message: "Invalid request payload",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	input := model.Category{
		Name: newCategory.Name,
		Description: newCategory.Description,
	}

	validate := validator.New()
	if err := validate.Struct(&input); err != nil {
        var errMsg strings.Builder; errMsg.WriteString("Validation failed: ")
        for _, fieldErr := range err.(validator.ValidationErrors) {
            errMsg .WriteString(fmt.Sprintf("%s (%s), ", fieldErr.Field(), fieldErr.Tag()))
        }
        
		errResponse := structs.ErrorResponse{
			Status: false,
			Message: errMsg.String(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
	}
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	return
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	return
}