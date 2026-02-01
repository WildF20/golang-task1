package product

import (
	"context"
	"golang-task1/internal/product/port"
)

type ProductService struct {
	repo *ProductRepository
	categoryChecker port.CategoryChecker
}

func NewProductService(categoryChecker port.CategoryChecker,repo *ProductRepository) *ProductService {
	return &ProductService{categoryChecker: categoryChecker, repo: repo}
}

func (s *ProductService) GetAll() ([]Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) Create(ctx context.Context, data *Product) error {
	exists, err := s.categoryChecker.ExistsByID(ctx, data.CategoryID)
    if err != nil {
        return err
    }

    if !exists {
        return ErrCategoryNotFound
    }

	return s.repo.Create(data)
}

func (s *ProductService) ExistsByID(id string) (bool, error) {
	return s.repo.ExistsByID(id)
}

func (s *ProductService) GetByID(id string) (*Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Update(product *Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) Delete(id string) error {
	return s.repo.Delete(id)
}
