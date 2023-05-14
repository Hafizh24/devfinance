package controller

import "github.com/hafizh24/devfinance/internal/app/schema"

type TransactionService interface {
	Create(req *schema.CreateTransactionReq) error
	ShowAll(req *schema.GetTransactionReq) ([]schema.GetTransactionResp, error)
	GetByType(req *schema.GetTransactionReq) ([]schema.GetTransactionResp, error)
	DeleteByID(id string) (*schema.GetTransactionResp, error)
	UpdateByID(id string, req *schema.UpdateTransactionReq) error
}
