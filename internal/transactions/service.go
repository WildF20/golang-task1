package transactions

type TransactionService struct {
	repo *TransactionRepository
}

func NewTransactionService(repo *TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}
func (s *TransactionService) Checkout(items []CheckoutItem) (*Transaction, error) {
		return s.repo.CreateTransaction(items)
}