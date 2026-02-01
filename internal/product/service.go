package product

type ProductService struct {
	repo *ProductRepository
}

func NewProductService(repo *ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll() ([]Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) Create(data *Product) error {
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
