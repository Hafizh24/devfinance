package model

import "time"

type Auth struct {
	ID       int       `db:"id"`
	Token    string    `db:"token"`
	AuthType string    `db:"auth_type"`
	UserID   int       `db:"user_id"`
	Expiry   time.Time `db:"expires_at"`
}
