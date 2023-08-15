package repo

import (
	"github.com/KlimGrishanov/bill-service/entity"
	"github.com/jmoiron/sqlx"
)

type Transaction interface {
	Create(transaction entity.Transaction) error
	Get(id int) entity.Transaction
}

type User interface {
	Create(user entity.User) error
	Get(id int) entity.User
	Update(id int, updateInfo entity.User) error
}

type Repository struct {
	Transaction
	User
}

func NewRepo(db *sqlx.DB) *Repository {
	return &Repository{
		Transaction: NewTransactionPostgres(db),
		User:        NewUserPostgres(db),
	}
}
