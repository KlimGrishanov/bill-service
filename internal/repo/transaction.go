package repo

import (
	"github.com/KlimGrishanov/bill-service/entity"
	"github.com/jmoiron/sqlx"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func (t TransactionPostgres) Create(transaction entity.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (t TransactionPostgres) Get(id int) entity.Transaction {
	//TODO implement me
	panic("implement me")
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}
