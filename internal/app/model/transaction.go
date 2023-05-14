package model

import "time"

type Transaction struct {
	ID          int       `db:"id"`
	Type        string    `db:"type"`
	Note        string    `db:"note"`
	Amount      int       `db:"amount"`
	CategoryID  int       `db:"category_id"`
	CurrencyID  int       `db:"currency_id"`
	UserID      int       `db:"user_id"`
	CreatedAt   time.Time `db:"created_at"`
	TotalAmount string    `db:"total_amount"`
	Category
	Currency
}
