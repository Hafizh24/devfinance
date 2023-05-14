package controller

import (
	"fmt"
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

func (tc *TransactionController) ShowRecaps(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.GetString("user_id"))
	req := &schema.GetTransactionReq{}

	req.UserID = id

	resp, err := tc.service.ShowAll(req)
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
	fmt.Println(req.UserID)

	err := tc.service.Create(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, "success create transaction", req)
}

func (tc *TransactionController) ShowByType(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.GetString("user_id"))
	param, _ := ctx.Params.Get("type")
	req := &schema.GetTransactionReq{}

	req.UserID = id
	req.Type = param

	resp, err := tc.service.GetByType(req)
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
