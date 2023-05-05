package service

import "github.com/hafizh24/devfinance/internal/app/model"

type CurrencyRepository interface {
	Browse() ([]model.Currency, error)
	Create(currency model.Currency) error
	GetByID(id string) (model.Currency, error)
	Update(id string, currency model.Currency) error
	Delete(id string) error
}
