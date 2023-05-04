package controller

import "github.com/hafizh24/devfinance/internal/app/schema"

type CategoryService interface {
	BrowseAll() ([]schema.GetCategoryResp, error)
	Create(req *schema.CreateCategoryReq) error
	GetByID(id string) (schema.GetCategoryResp, error)
	UpdateByID(id string, req *schema.UpdateCategoryReq) error
	DeleteByID(id string) (*schema.GetCategoryResp, error)
}
