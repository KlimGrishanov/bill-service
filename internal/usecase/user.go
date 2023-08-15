package usecase

import (
	"github.com/KlimGrishanov/bill-service/entity"
	"github.com/KlimGrishanov/bill-service/internal/repo"
)

type UserUseCase struct {
	repo repo.User
}

func (u UserUseCase) Create(user entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserUseCase) Get(id int) entity.User {
	//TODO implement me
	panic("implement me")
}

func (u UserUseCase) Update(id int, updateInfo entity.User) error {
	//TODO implement me
	panic("implement me")
}

func NewUserUseCase(repo repo.User) *UserUseCase {
	return &UserUseCase{repo: repo}
}
