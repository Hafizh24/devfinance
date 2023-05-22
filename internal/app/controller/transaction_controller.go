package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hafizh24/devfinance/internal/app/schema"
	"github.com/hafizh24/devfinance/internal/pkg/handler"
)

type TransactionController struct {
	service TransactionService
}

func NewTransactionController(service TransactionService) *TransactionController {
	return &TransactionController{service: service}
}

func (tc *TransactionController) BrowseTransaction(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.GetString("user_id"))
	req := &schema.BrowseTransactionReq{}

	req.Page = ctx.GetInt("page")
	req.PageSize = ctx.GetInt("page_size")
	req.Type = ctx.GetString("type")
	req.UserID = id

	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp, err := tc.service.BrowseAll(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success get list transaction", resp)
}

func (tc *TransactionController) CreateTransaction(ctx *gin.Context) {
	req := &schema.CreateTransactionReq{}
	id, _ := strconv.Atoi(ctx.GetString("user_id"))

	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.UserID = id

	err := tc.service.Create(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, "success create transaction", req)
}

func (tc *TransactionController) DetailTransaction(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	resp, err := tc.service.GetByID(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success get detail transaction", resp)
}

func (tc *TransactionController) DeleteTransaction(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	_, err := tc.service.DeleteByID(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success deleted transaction", nil)
}

func (tc *TransactionController) UpdateTransaction(ctx *gin.Context) {
	req := &schema.UpdateTransactionReq{}
	postID, _ := ctx.Params.Get("id")

	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := tc.service.UpdateByID(postID, req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success update transaction", nil)
}
