package schema

import "time"

type GetCategoryResp struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateCategoryReq struct {
	Name        string    `validate:"required" json:"name"`
	Description string    `validate:"required" json:"description"`
	Created_at  time.Time `json:"created_at"`
}

type UpdateCategoryReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
