package category

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang-task1/internal/shared/structs"
)

var (
	validate   = validator.New()
)

type CategoryHandler struct {
	service *CategoryService
}

func NewCategoryHandler(service *CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAll()
	if err != nil {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errResponse)
		return
	}
	
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

func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	if err := validate.Var(idStr, "required,uuid"); err != nil {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: "Invalid ID format",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	category, err := h.service.GetByID(idStr)
	if err != nil {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: err.Error(),
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

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
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

	if err := h.service.Create(&newCategory); err != nil {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	response := structs.SuccessResponse{
		Status:  true,
		Message: "Category created successfully",
		Data:    newCategory,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	var payload Category
	idStr := r.PathValue("id")
	ctx := r.Context()

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

	if err := validate.Var(idStr, "required,uuid"); err != nil {
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

	categoryExists, err := h.service.ExistsByID(ctx, idStr)
	if err != nil {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	if !categoryExists {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: "Category not found",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	updatedCategory := &Category{
		ID:          idStr,
		Name:        payload.Name,
		Description: payload.Description,
	}

	if err := h.service.Update(updatedCategory); err != nil {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	response := structs.SuccessResponse{
		Status:  true,
		Message: "Category updated successfully",
		Data:    updatedCategory,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	ctx := r.Context()

	if err := validate.Var(idStr, "required,uuid"); err != nil {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: "Invalid ID format",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	categoryExists, err := h.service.ExistsByID(ctx, idStr)
	if err != nil {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	if !categoryExists {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: "Category not found",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	if err := h.service.Delete(idStr); err != nil {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	response := structs.SuccessResponse{
		Status:  true,
		Message: "Category deleted successfully",
		Data: nil,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
