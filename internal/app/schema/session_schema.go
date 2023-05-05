package schema

type LoginReq struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required,min=8,alphanum" json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResp struct {
	AccessToken string `json:"access_token"`
}

type RefreshTokenReq struct {
	UserID       int    `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
}
type LogoutReq struct {
	UserID int `json:"user_id"`
}

type ShowReq struct {
	UserID int `json:"user_id"`
}
type ShowResp struct {
	Fullname  string `json:"fullname"`
	UserSince string `json:"user_since"`
	Username  string `validate:"required" json:"username"`
}
