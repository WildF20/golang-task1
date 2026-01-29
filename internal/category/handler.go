package category

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/oklog/ulid/v2"
	"golang-task1/internal/shared/structs"
)

var (
	categories = make(map[string]*Category)
	mu         = sync.RWMutex{}
	validate   = validator.New()
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	response := structs.SuccessResponse{
		Status:  true,
		Message: "Success",
		Data:    categories,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errResponse)
		return
	}
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	if err := validate.Var(idStr, "required,ulid"); err != nil {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: "Invalid ID format",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	mu.RLock()
	defer mu.RUnlock()

	category, exists := categories[idStr]
	if !exists {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: "Category not found",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	response := structs.SuccessResponse{
		Status:  true,
		Message: "Success",
		Data:    category,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var payload Category

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	newCategory := Category{
		Name:        payload.Name,
		Description: payload.Description,
	}

	if err := validate.Struct(&payload); err != nil {
		var errMsg strings.Builder
		errMsg.WriteString("Validation failed: ")
		for _, fieldErr := range err.(validator.ValidationErrors) {
			errMsg.WriteString(fmt.Sprintf("%s (%s), ", fieldErr.Field(), fieldErr.Tag()))
		}

		errResponse := structs.ErrorResponse{
			Status:  false,
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

	categories[newCategory.ID] = &newCategory

	response := structs.SuccessResponse{
		Status:  true,
		Message: "Category created successfully",
		Data:    newCategory,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var payload Category
	idStr := r.PathValue("id")

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	if err := validate.Var(idStr, "required,ulid"); err != nil {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: "Invalid ID format",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	if err := validate.Struct(&payload); err != nil {
		var errMsg strings.Builder
		errMsg.WriteString("Validation failed: ")
		for _, fieldErr := range err.(validator.ValidationErrors) {
			errMsg.WriteString(fmt.Sprintf("%s (%s), ", fieldErr.Field(), fieldErr.Tag()))
		}

		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: errMsg.String(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	category, exists := categories[idStr]
	if !exists {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: "Category not found",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	category.Name = payload.Name
	category.Description = payload.Description

	response := structs.SuccessResponse{
		Status:  true,
		Message: "Category updated successfully",
		Data:    categories[idStr],
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	if err := validate.Var(idStr, "required,ulid"); err != nil {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: "Invalid ID format",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	_, exists := categories[idStr]
	if !exists {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: "Category not found",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	delete(categories, idStr)

	response := structs.SuccessResponse{
		Status:  true,
		Message: "Category deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
