package service

import (
	"errors"

	"github.com/hafizh24/devfinance/internal/app/model"
	"github.com/hafizh24/devfinance/internal/app/schema"
	"github.com/hafizh24/devfinance/internal/pkg/reason"
	"golang.org/x/crypto/bcrypt"
)

type RegistrationService struct {
	repo UserRepository
}

func NewRegistrationService(repo UserRepository) *RegistrationService {
	return &RegistrationService{repo: repo}
}

func (rs *RegistrationService) Register(req *schema.RegisterReq) error {

	existingUser, _ := rs.repo.GetByUsername(req.Username)
	if existingUser.ID > 0 {
		return errors.New(reason.UserAlreadyExist)
	}

	password, _ := rs.hashPassword(req.Password)

	var insertData model.User
	insertData.Fullname = req.Fullname
	insertData.Password = password
	insertData.Username = req.Username

	err := rs.repo.Create(insertData)
	if err != nil {
		return errors.New(reason.RegisterFailed)
	}

	return nil
}

func (rs *RegistrationService) hashPassword(password string) (string, error) {
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(bytePassword), nil

}

func (rs *RegistrationService) DeleteByID(id int) (*schema.GetUserResp, error) {
	resp := &schema.GetUserResp{}

	check, err := rs.repo.GetByID(id)
	if check.ID == 0 && err != nil {
		return nil, errors.New(reason.UserNotFound)
	}

	user, err := rs.repo.Delete(id)
	if err != nil {
		return resp, errors.New(reason.UserCannotDelete)
	}

	resp.Fullname = user.Fullname
	resp.Username = user.Username
	resp.Password = user.Password

	return resp, nil
}
