package schema

import "mime/multipart"

type GetTransactionResp struct {
	ID       int     `json:"id"`
	Date     string  `json:"date"`
	Amount   string  `json:"amount"`
	Category string  `json:"category"`
	Type     string  `json:"type"`
	Note     string  `json:"note"`
	ImageUrl *string `json:"image_url"`
}
type CreateTransactionReq struct {
	Type       string                `validate:"required,oneof=expenses income" form:"type"`
	Note       string                `validate:"required" form:"note"`
	Amount     int                   `validate:"required" form:"amount"`
	CategoryID int                   `validate:"required" form:"category_id"`
	CurrencyID int                   `validate:"required" form:"currency_id"`
	Image      *multipart.FileHeader `validate:"required" form:"image"`
	UserID     int                   `json:"user_id"`
}

type UpdateTransactionReq struct {
	Type       string `validate:"required,oneof=expenses income" json:"type"`
	Note       string `validate:"required" json:"note"`
	Amount     int    `validate:"required" json:"amount"`
	CategoryID int    `validate:"required" json:"category_id"`
	CurrencyID int    `validate:"required" json:"currency_id"`
}

type BrowseTransactionReq struct {
	Page     int
	PageSize int
	UserID   int    `json:"user_id"`
	Type     string `validate:"omitempty,oneof=expenses income" json:"type"`
}
