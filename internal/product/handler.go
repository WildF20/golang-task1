package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"golang-task1/internal/shared/structs"

	"github.com/go-playground/validator/v10"
)

var (
	validate   = validator.New()
)

type ProductHandler struct {
	service *ProductService
}

func NewProductHandler(service *ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	products, err := h.service.GetAll(name)
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
		Data:    products,
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

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	product, err := h.service.GetByID(idStr)
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
		Data:    product,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var payload Product
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

	newProduct := Product{
		Name:        payload.Name,
		CategoryID:  payload.CategoryID,
		Price:       payload.Price,
		Stock:       payload.Stock,
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

	if err := h.service.Create(ctx, &newProduct); err != nil {
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
		Message: "Product created successfully",
		Data:    newProduct,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	var payload Product
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

	productExists, err := h.service.ExistsByID(idStr)
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

	if !productExists {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: "Product not found",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	updatedProduct := &Product{
		ID:          idStr,
		Name:        payload.Name,
		CategoryID:  payload.CategoryID,
		Price:       payload.Price,
		Stock:       payload.Stock,
	}

	if err := h.service.Update(updatedProduct); err != nil {
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
		Message: "Product updated successfully",
		Data:    updatedProduct,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	productExists, err := h.service.ExistsByID(idStr)
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

	if !productExists {
		errResponse := structs.ErrorResponse{
			Status:  false,
			Message: "Product not found",
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
		Message: "Product deleted successfully",
		Data: nil,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
