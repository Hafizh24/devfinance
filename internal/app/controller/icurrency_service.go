package controller

import "github.com/hafizh24/devfinance/internal/app/schema"

type CurrencyService interface {
	BrowseAll() ([]schema.GetCurrencyResp, error)
	Create(req *schema.CreateCurrencyReq) error
	GetByID(id string) (schema.GetCurrencyResp, error)
	UpdateByID(id string, req *schema.UpdateCurrencyReq) error
	DeleteByID(id string) (*schema.GetCurrencyResp, error)
}
