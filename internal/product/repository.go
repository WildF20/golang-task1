package product

import (
	"database/sql"
	"errors"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll() ([]Product, error) {
	query := "SELECT id, name, category_id, price, stock FROM products"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]Product, 0)
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.CategoryID, &p.Price, &p.Stock)
		if err != nil {
			return nil, err
		}
		categories = append(categories, p)
	}

	return categories, nil
}

func (repo *ProductRepository) ExistsByID(id string) (bool, error) {
	query := "SELECT COUNT(1) FROM products WHERE id = $1"
	var count int
	err := repo.db.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (repo *ProductRepository) GetByID(id string) (*Product, error) {
	query := "SELECT id, name, category_id, price, stock FROM products WHERE id = $1"

	var p Product
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.CategoryID, &p.Price, &p.Stock)
	if err == sql.ErrNoRows {
		return nil, errors.New("product tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}


func (repo *ProductRepository) Create(product *Product) error {
	query := "INSERT INTO products (name, category_id, price, stock) VALUES ($1, $2, $3, $4) RETURNING id"
	err := repo.db.QueryRow(query, product.Name, product.CategoryID, product.Price, product.Stock).Scan(&product.ID)
	return err
}

func (repo *ProductRepository) Update(product *Product) error {
	query := "UPDATE products SET name = $1, category_id = $2, price = $3, stock = $4 WHERE id = $5"
	result, err := repo.db.Exec(query, product.Name, product.CategoryID, product.Price, product.Stock, product.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product tidak ditemukan")
	}

	return nil
}

func (repo *ProductRepository) Delete(id string) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product tidak ditemukan")
	}

	return err
}
