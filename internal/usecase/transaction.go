package usecase

import (
	"github.com/KlimGrishanov/bill-service/entity"
	"github.com/KlimGrishanov/bill-service/internal/repo"
)

type TransactionUseCase struct {
	repo repo.Transaction
}

func (t TransactionUseCase) Create(transaction entity.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (t TransactionUseCase) Get(id int) entity.Transaction {
	//TODO implement me
	panic("implement me")
}

func NewTransactionUseCase(repo repo.Transaction) *TransactionUseCase {
	return &TransactionUseCase{repo: repo}
}
