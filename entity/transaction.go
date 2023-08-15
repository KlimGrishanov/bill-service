package entity

type Transaction struct {
	Id              uint   `json:"id" db:"id"`
	IdFrom          uint   `json:"idFrom" db:"id_from"`
	IdTo            uint   `json:"idTo" db:"id_to"`
	Amount          uint   `json:"amount" db:"amount"`
	CurrencyType    string `json:"currencyType" db:"currency_type"`
	TransactionType string `json:"transactionType" db:"transaction_type"`
}
