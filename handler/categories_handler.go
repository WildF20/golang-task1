package handler

import (
	"strings"
	"fmt"
	"sync"
	"net/http"
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/oklog/ulid/v2"
	
	"golang-task1/model"
	"golang-task1/structs"
)

var (
	categories 	= []model.Category{}
	mu         	= sync.RWMutex{}
	validate 	= validator.New()
)

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
    defer mu.RUnlock()

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
		return
	}
}

func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	return
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var payload model.Category

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		errResponse := structs.ErrorResponse{
			Status: false,
			Message: err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	newCategory := model.Category{
		Name: payload.Name,
		Description: payload.Description,
	}

	if err := validate.Struct(&payload); err != nil {
        var errMsg strings.Builder; errMsg.WriteString("Validation failed: ")
        for _, fieldErr := range err.(validator.ValidationErrors) {
            errMsg.WriteString(fmt.Sprintf("%s (%s), ", fieldErr.Field(), fieldErr.Tag()))
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

	mu.Lock()
	defer mu.Unlock()

	id := ulid.Make().String()
	newCategory.ID = id

	categories = append(categories, newCategory)

	response := structs.SuccessResponse{
		Status: true,
		Message: "Category created successfully",
		Data: newCategory,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	return
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	return
}