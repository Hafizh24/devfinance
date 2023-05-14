package schema

type RegisterReq struct {
	Fullname string `validate:"required" json:"fullname"`
	Password string `validate:"required,min=8,alphanum" json:"password"`
	Username string `validate:"required,alphanum" json:"username"`
}

type GetUserResp struct {
	Fullname string `json:"fullname"`
	Password string `json:"password"`
	Username string `json:"username"`
}
