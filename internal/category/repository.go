package category

import "github.com/google/uuid"

type Repository interface {
	GetAll() []Category
	Create(p Category) Category
}

type memoryRepository struct {
	data map[string]Category
}

func NewRepository() Repository {
	return &memoryRepository{
		data: make(map[string]Category),
	}
}

func (r *memoryRepository) GetAll() []Category {
	result := make([]Category, 0, len(r.data))
	for _, p := range r.data {
		result = append(result, p)
	}
	return result
}

func (r *memoryRepository) Create(p Category) Category {
	p.ID = uuid.NewString()
	r.data[p.ID] = p
	return p
}