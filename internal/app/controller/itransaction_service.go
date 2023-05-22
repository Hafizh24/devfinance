package controller

import "github.com/hafizh24/devfinance/internal/app/schema"

type TransactionService interface {
	Create(req *schema.CreateTransactionReq) error
	BrowseAll(req *schema.BrowseTransactionReq) ([]schema.GetTransactionResp, error)
	GetByID(id string) (schema.GetTransactionResp, error)
	DeleteByID(id string) (*schema.GetTransactionResp, error)
	UpdateByID(id string, req *schema.UpdateTransactionReq) error
}
