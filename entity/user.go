package entity

type User struct {
	Id      uint    `json:"id" db:"id"`
	Balance float64 `json:"balance" db:"balance"`
}
