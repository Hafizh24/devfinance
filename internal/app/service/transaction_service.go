package service

import (
	"errors"
	"fmt"

	"mime/multipart"

	"github.com/hafizh24/devfinance/internal/app/model"
	"github.com/hafizh24/devfinance/internal/app/schema"
	"github.com/hafizh24/devfinance/internal/pkg/reason"
	log "github.com/sirupsen/logrus"
)

type ImageUploader interface {
	UploadImage(input *multipart.FileHeader) (string, error)
}

type TransactionService struct {
	transactionrepo TransactionRepository
	authrepo        AuthRepository
	uploader        ImageUploader
}

func NewTransactionService(transactionrepo TransactionRepository, authrepo AuthRepository, uploader ImageUploader) *TransactionService {
	return &TransactionService{transactionrepo: transactionrepo, authrepo: authrepo, uploader: uploader}
}

func (ts *TransactionService) Create(req *schema.CreateTransactionReq) error {
	var insertData model.Transaction

	insertData.Type = req.Type
	insertData.Note = req.Note
	insertData.Amount = req.Amount
	insertData.CategoryID = req.CategoryID
	insertData.CurrencyID = req.CurrencyID
	insertData.UserID = req.UserID

	imageURL, err := ts.uploader.UploadImage(req.Image)
	if err != nil {
		log.Error("upload image transaction %w", err)
		return errors.New(reason.TransactionCannotCreate)
	}

	insertData.ImageUrl = &imageURL

	err = ts.transactionrepo.Create(insertData)
	if err != nil {
		return errors.New(reason.TransactionCannotCreate)
	}
	return nil
}

func (ts *TransactionService) ShowAll(req *schema.GetTransactionReq) ([]schema.GetTransactionResp, error) {
	var resp []schema.GetTransactionResp

	existingUser, _ := ts.authrepo.GetByUserID(req.UserID)
	if existingUser.ID <= 0 {
		return nil, errors.New(reason.UserNotFound)
	}

	transactions, err := ts.transactionrepo.Browse(req.UserID)
	if err != nil {
		return nil, errors.New(reason.TransactionCannotBrowse)
	}

	for _, value := range transactions {
		var respData schema.GetTransactionResp
		respData.ID = value.ID
		respData.Date = value.CreatedAt.Format("02-01-2006")
		respData.Amount = value.TotalAmount
		respData.Category = value.Name
		respData.Type = value.Type
		respData.Note = value.Note
		respData.ImageUrl = value.ImageUrl
		resp = append(resp, respData)
	}

	return resp, nil
}

func (ts *TransactionService) GetByType(req *schema.GetTransactionReq) ([]schema.GetTransactionResp, error) {
	var resp []schema.GetTransactionResp
	// var respData schema.GetTransactionResp

	fmt.Println(req.UserID)

	transactions, err := ts.transactionrepo.GetByType(req.Type, req.UserID)
	if err != nil {
		return nil, errors.New(reason.TransactionCannotGetDetail)
	}

	for _, value := range transactions {
		var respData schema.GetTransactionResp
		respData.ID = value.ID
		respData.Date = value.CreatedAt.Format("02-01-2006")
		respData.Amount = value.TotalAmount
		respData.Category = value.Name
		respData.Type = value.Type
		respData.Note = value.Note
		respData.ImageUrl = value.ImageUrl
		resp = append(resp, respData)
	}

	return resp, nil
}

func (ts *TransactionService) UpdateByID(id string, req *schema.UpdateTransactionReq) error {
	var updateData model.Transaction

	updateData.Type = req.Type
	updateData.Note = req.Note
	updateData.Amount = req.Amount
	updateData.CategoryID = req.CategoryID
	updateData.CurrencyID = req.CurrencyID

	check, err := ts.transactionrepo.GetByID(id)
	if check.ID == 0 && err != nil {
		return errors.New(reason.TransactionNotFound)
	}

	err = ts.transactionrepo.Update(id, updateData)
	if err != nil {
		return errors.New(reason.TransactionCannotUpdate)
	}

	return nil
}

func (ts *TransactionService) DeleteByID(id string) (*schema.GetTransactionResp, error) {
	resp := &schema.GetTransactionResp{}

	check, _ := ts.transactionrepo.GetByID(id)
	if check.ID == 0 {
		return nil, errors.New(reason.TransactionNotFound)
	}

	err := ts.transactionrepo.Delete(id)
	if err != nil {
		return resp, errors.New(reason.TransactionCannotDelete)
	}

	return resp, nil
}

func (ts *TransactionService) GetByID(id string) (schema.GetTransactionResp, error) {
	var resp schema.GetTransactionResp

	transaction, err := ts.transactionrepo.GetByID(id)
	if err != nil {
		return resp, errors.New(reason.TransactionCannotGetDetail)
	}

	resp.ID = transaction.ID
	resp.Date = transaction.CreatedAt.Format("02-01-2006")
	resp.Amount = transaction.TotalAmount
	resp.Category = transaction.Name
	resp.Type = transaction.Type
	resp.Note = transaction.Note
	resp.ImageUrl = transaction.ImageUrl

	return resp, nil
}
