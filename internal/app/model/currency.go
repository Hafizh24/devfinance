package model

type Currency struct {
	ID          int    `db:"id"`
	Code        string `db:"code"`
	Description string `db:"description"`
}
