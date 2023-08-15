package usecase

import (
	"github.com/KlimGrishanov/bill-service/entity"
	"github.com/KlimGrishanov/bill-service/internal/repo"
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

type UseCase struct {
	Transaction
	User
}

func NewUseCase(repo *repo.Repository) *UseCase {
	return &UseCase{
		Transaction: NewTransactionUseCase(repo.Transaction),
		User:        NewUserUseCase(repo.User),
	}
}
