package service

import "github.com/hafizh24/devfinance/internal/app/model"

type CategoryRepository interface {
	Browse() ([]model.Category, error)
	Create(category model.Category) error
	GetByID(id string) (model.Category, error)
	Update(id string, category model.Category) error
	Delete(id string) (model.Category, error)
}
