package service

import "github.com/hafizh24/devfinance/internal/app/model"

type TransactionRepository interface {
	Browse(search model.BrowseTransaction) ([]model.Transaction, error)
	Create(transaction model.Transaction) error
	Update(id string, transaction model.Transaction) error
	Delete(id string) error
	GetByID(id string) (model.Transaction, error)
}
