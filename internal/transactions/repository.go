package transactions

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(items []CheckoutItem) (*Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]TransactionDetail, 0)

	for _, item := range items {
		var productPrice, stock int
		var productName string

		err := tx.QueryRow("SELECT name, price, stock FROM products WHERE id = $1", item.ProductID).Scan(&productName, &productPrice, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %s not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}
  
		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			log.Println("Error updating product stock:", err)
			return nil, err
		}

		details = append(details, TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id", totalAmount).Scan(&transactionID)
	if err != nil {
		log.Println("Error inserting transaction:", err)
		return nil, err
	}

	values := []interface{}{}
	query := "INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES "

	placeholders := []string{}

	for i, d := range details {
		start := i*4 + 1

		placeholders = append(placeholders,
			fmt.Sprintf("($%d,$%d,$%d,$%d)", start, start+1, start+2, start+3),
		)

		values = append(values,
			transactionID,
			d.ProductID,
			d.Quantity,
			d.Subtotal,
		)
	}

	query += strings.Join(placeholders, ",")
	_, err = tx.Exec(query, values...)
	
	if err != nil {
		log.Println("Error inserting transaction detail:", err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Println("Error committing transaction:", err)
		return nil, err
	}

	return &Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}