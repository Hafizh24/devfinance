package repository

import (
	"fmt"

	"github.com/hafizh24/devfinance/internal/app/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type AuthRepository struct {
	DB *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (ar *AuthRepository) Create(auth model.Auth) error {
	var sqlStatement = `
		INSERT INTO auths (token, auth_type, user_id, expires_at)
		VALUES ($1, $2, $3, $4)
		`

	_, err := ar.DB.Exec(sqlStatement, auth.Token, auth.AuthType, auth.UserID, auth.Expiry)
	if err != nil {
		log.Error(fmt.Errorf("error AuthRepository - Create : %w", err))
		return err
	}
	return nil
}
func (ar *AuthRepository) Find(userID int, RefreshToken string) (model.Auth, error) {
	var auth model.Auth
	var sqlStatement = `
		SELECT id, token, auth_type, user_id, expires_at
		FROM auths
		WHERE user_id = $1 AND token = $2
		`

	err := ar.DB.QueryRowx(sqlStatement, userID, RefreshToken).StructScan(&auth)
	if err != nil {
		log.Error(fmt.Errorf("error AuthRepository - Find : %w", err))
		return auth, err
	}

	return auth, nil
}
func (ar *AuthRepository) Delete(userID int) error {

	var sqlStatement = `
	DELETE FROM auths
	WHERE user_id = $1
		`

	_, err := ar.DB.Exec(sqlStatement, userID)
	if err != nil {
		log.Error(fmt.Errorf("error AuthRepository - Delete : %w", err))
		return err
	}

	return nil
}
