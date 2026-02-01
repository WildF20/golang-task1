package category

import "context"

type CategoryService struct {
	repo *CategoryRepository
}

func NewCategoryService(repo *CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll() ([]Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) Create(data *Category) error {
	return s.repo.Create(data)
}

func (s *CategoryService) ExistsByID(ctx context.Context, id string) (bool, error) {
	return s.repo.ExistsByID(ctx, id)
}

func (s *CategoryService) GetByID(id string) (*Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) Update(product *Category) error {
	return s.repo.Update(product)
}

func (s *CategoryService) Delete(id string) error {
	return s.repo.Delete(id)
}
