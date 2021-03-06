package league

import (
	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/transaction"
)

// Repository interface
type Repository interface {
	GetByID(tr transaction.Transaction, id string) (entity.League, error)
	Add(tr transaction.Transaction, league entity.League) (entity.League, error)
	Update(tr transaction.Transaction, league entity.League) error
}
