package schema

type GetCurrencyResp struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type CreateCurrencyReq struct {
	Code        string `validate:"required" json:"code"`
	Description string `validate:"required" json:"description"`
}

type UpdateCurrencyReq struct {
	Code        string `json:"code" validate:"required"`
	Description string `json:"description" validate:"required"`
}
