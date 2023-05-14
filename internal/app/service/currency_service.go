package service

import (
	"errors"

	"github.com/hafizh24/devfinance/internal/app/model"
	"github.com/hafizh24/devfinance/internal/app/schema"
	"github.com/hafizh24/devfinance/internal/pkg/reason"
)

type CurrencyService struct {
	repo CurrencyRepository
}

func NewCurrencyService(repo CurrencyRepository) *CurrencyService {
	return &CurrencyService{repo: repo}
}

func (cs *CurrencyService) Create(req *schema.CreateCurrencyReq) error {
	var insertData model.Currency

	insertData.Code = req.Code
	insertData.Description = req.Description

	err := cs.repo.Create(insertData)
	if err != nil {
		return errors.New(reason.CurrencyCannotCreate)
	}
	return nil
}

func (cs *CurrencyService) BrowseAll() ([]schema.GetCurrencyResp, error) {
	var resp []schema.GetCurrencyResp

	categories, err := cs.repo.Browse()
	if err != nil {
		return nil, errors.New(reason.CurrencyCannotBrowse)
	}

	for _, value := range categories {
		var respData schema.GetCurrencyResp
		respData.ID = value.ID
		respData.Code = value.Code
		respData.Description = value.Description
		resp = append(resp, respData)
	}

	return resp, nil
}

func (cs *CurrencyService) GetByID(id string) (schema.GetCurrencyResp, error) {
	var resp schema.GetCurrencyResp

	Currency, err := cs.repo.GetByID(id)
	if err != nil {
		return resp, errors.New(reason.CurrencyCannotGetDetail)
	}

	resp.ID = Currency.ID
	resp.Code = Currency.Code
	resp.Description = Currency.Description

	return resp, nil
}

func (cs *CurrencyService) UpdateByID(id string, req *schema.UpdateCurrencyReq) error {
	var updateData model.Currency

	updateData.Code = req.Code
	updateData.Description = req.Description

	check, err := cs.repo.GetByID(id)
	if check.ID == 0 && err != nil {
		return errors.New(reason.CurrencyNotFound)
	}

	err = cs.repo.Update(id, updateData)
	if err != nil {
		return errors.New(reason.CurrencyCannotUpdate)
	}

	return nil
}

func (cs *CurrencyService) DeleteByID(id string) (*schema.GetCurrencyResp, error) {
	resp := &schema.GetCurrencyResp{}

	check, _ := cs.repo.GetByID(id)
	if check.ID == 0 {
		return nil, errors.New(reason.CurrencyNotFound)
	}

	err := cs.repo.Delete(id)
	if err != nil {
		return resp, errors.New(reason.CurrencyCannotDelete)
	}

	return resp, nil
}
