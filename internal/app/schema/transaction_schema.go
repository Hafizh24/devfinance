package schema

type GetTransactionResp struct {
	ID       int    `json:"id"`
	Date     string `json:"date"`
	Amount   string `json:"amount"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Note     string `json:"note"`
}
type GetTransactionReq struct {
	UserID int    `json:"user_id"`
	Type   string `json:"type"`
}

type CreateTransactionReq struct {
	Type       string `validate:"required,oneof=expenses income" json:"type"`
	Note       string `validate:"required" json:"note"`
	Amount     int    `validate:"required" json:"amount"`
	CategoryID int    `validate:"required" json:"category_id"`
	CurrencyID int    `validate:"required" json:"currency_id"`
	UserID     int    `json:"user_id"`
}

type UpdateTransactionReq struct {
	Type       string `validate:"required,oneof=expenses income" json:"type"`
	Note       string `validate:"required" json:"note"`
	Amount     int    `validate:"required" json:"amount"`
	CategoryID int    `validate:"required" json:"category_id"`
	CurrencyID int    `validate:"required" json:"currency_id"`
}
