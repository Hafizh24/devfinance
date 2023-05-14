package service

import (
	"errors"

	"github.com/hafizh24/devfinance/internal/app/model"
	"github.com/hafizh24/devfinance/internal/app/schema"
	"github.com/hafizh24/devfinance/internal/pkg/reason"
)

type CategoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (cs *CategoryService) Create(req *schema.CreateCategoryReq) error {
	var insertData model.Category

	insertData.Name = req.Name
	insertData.Description = req.Description

	err := cs.repo.Create(insertData)
	if err != nil {
		return errors.New(reason.CategoryCannotCreate)
	}
	return nil
}

func (cs *CategoryService) BrowseAll() ([]schema.GetCategoryResp, error) {
	var resp []schema.GetCategoryResp

	categories, err := cs.repo.Browse()
	if err != nil {
		return nil, errors.New(reason.CategoryCannotBrowse)
	}

	for _, value := range categories {
		var respData schema.GetCategoryResp
		respData.ID = value.ID
		respData.Name = value.Name
		respData.Description = value.Description
		resp = append(resp, respData)
	}

	return resp, nil
}

func (cs *CategoryService) GetByID(id string) (schema.GetCategoryResp, error) {
	var resp schema.GetCategoryResp

	category, err := cs.repo.GetByID(id)
	if err != nil {
		return resp, errors.New(reason.CategoryCannotGetDetail)
	}

	resp.ID = category.ID
	resp.Name = category.Name
	resp.Description = category.Description

	return resp, nil
}

func (cs *CategoryService) UpdateByID(id string, req *schema.UpdateCategoryReq) error {
	var updateData model.Category

	updateData.Name = req.Name
	updateData.Description = req.Description

	check, err := cs.repo.GetByID(id)
	if check.ID == 0 && err != nil {
		return errors.New(reason.CategoryNotFound)
	}

	err = cs.repo.Update(id, updateData)
	if err != nil {
		return errors.New(reason.CategoryCannotUpdate)
	}

	return nil
}

func (cs *CategoryService) DeleteByID(id string) (*schema.GetCategoryResp, error) {

	resp := &schema.GetCategoryResp{}

	check, err := cs.repo.GetByID(id)
	if check.ID == 0 && err != nil {
		return nil, errors.New(reason.CategoryNotFound)
	}

	category, err := cs.repo.Delete(id)
	if err != nil {
		return resp, errors.New(reason.CategoryCannotDelete)
	}

	resp.ID = category.ID
	resp.Name = category.Name
	resp.Description = category.Description

	return resp, nil
}
