package service

import "github.com/hafizh24/devfinance/internal/app/model"

type UserRepository interface {
	Create(user model.User) error
	GetByUsername(username string) (model.User, error)
	GetByID(id int) (model.User, error)
}
