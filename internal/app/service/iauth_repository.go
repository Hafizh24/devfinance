package service

import "github.com/hafizh24/devfinance/internal/app/model"

type AuthRepository interface {
	Create(auth model.Auth) error
	Find(userID int, RefreshToken string) (model.Auth, error)
	Delete(userID int) error
}
