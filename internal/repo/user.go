package repo

import (
	"github.com/KlimGrishanov/bill-service/entity"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func (u UserPostgres) Create(user entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserPostgres) Get(id int) entity.User {
	//TODO implement me
	panic("implement me")
}

func (u UserPostgres) Update(id int, updateInfo entity.User) error {
	//TODO implement me
	panic("implement me")
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}
