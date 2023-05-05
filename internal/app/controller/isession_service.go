package controller

import "github.com/hafizh24/devfinance/internal/app/schema"

type SessionService interface {
	Login(req *schema.LoginReq) (schema.LoginResp, error)
	RefreshToken(req *schema.RefreshTokenReq) (schema.RefreshTokenResp, error)
	Logout(req *schema.LogoutReq) error
}
