package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hafizh24/devfinance/internal/app/schema"
	"github.com/hafizh24/devfinance/internal/pkg/handler"
)

type CurrencyController struct {
	service CurrencyService
}

func NewCurrencyController(service CurrencyService) *CurrencyController {
	return &CurrencyController{service: service}
}

func (cc *CurrencyController) BrowseCurrency(ctx *gin.Context) {
	resp, err := cc.service.BrowseAll()
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success get list currency", resp)
}

func (cc *CurrencyController) CreateCurrency(ctx *gin.Context) {
	req := &schema.CreateCurrencyReq{}

	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := cc.service.Create(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, "success create currency", req)
}

func (cc *CurrencyController) DetailCurrency(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	resp, err := cc.service.GetByID(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success get detail currency", resp)
}

func (cc *CurrencyController) UpdateCurrency(ctx *gin.Context) {
	req := &schema.UpdateCurrencyReq{}
	postID, _ := ctx.Params.Get("id")

	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := cc.service.UpdateByID(postID, req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success update currency", nil)
}

func (cc *CurrencyController) DeleteCurrency(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	_, err := cc.service.DeleteByID(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success deleted currency", nil)
}
