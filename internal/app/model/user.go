package model

import "time"

type User struct {
	ID        int       `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Fullname  string    `db:"fullname"`
	CreatedAt time.Time `db:"created_at"`
}
