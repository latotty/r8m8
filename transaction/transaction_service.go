package transaction

// Service interface
type Service interface {
	Start() (Transaction, error)
	Commit(transaction Transaction) error
	Rollback(transaction Transaction) error
	CommitOrRollback(transaction Transaction)
}
