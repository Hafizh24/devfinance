package repository

import (
	"fmt"

	"github.com/hafizh24/devfinance/internal/app/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type CurrencyRepository struct {
	DB *sqlx.DB
}

func NewCurrencyRepository(db *sqlx.DB) *CurrencyRepository {
	return &CurrencyRepository{DB: db}
}

func (cr *CurrencyRepository) Create(currency model.Currency) error {
	var sqlStatement = `
		INSERT INTO currencies (code, description)
		VALUES ($1, $2)
		`

	_, err := cr.DB.Exec(sqlStatement, currency.Code, currency.Description)
	if err != nil {
		log.Error(fmt.Errorf("error CurrencyRepository - Create : %w", err))
		return err
	}
	return nil
}

func (cr *CurrencyRepository) Browse() ([]model.Currency, error) {

	var currencies []model.Currency
	var sqlStatement = `
		SELECT id,code,description
		FROM currencies
	`
	rows, err := cr.DB.Queryx(sqlStatement)
	if err != nil {
		log.Error(fmt.Errorf("error CurrencyRepository - Browse : %w", err))
		return currencies, err
	}

	for rows.Next() {
		var currency model.Currency
		// nolint:errcheck
		rows.StructScan(&currency)
		currencies = append(currencies, currency)
	}
	return currencies, nil
}

func (cr *CurrencyRepository) GetByID(id string) (model.Currency, error) {
	var currency model.Currency
	var sqlStatement = `
	SELECT id,code,description
	FROM currencies
	WHERE id = $1
	`
	err := cr.DB.QueryRowx(sqlStatement, id).StructScan(&currency)
	if err != nil {
		log.Error(fmt.Errorf("error CurrencyRepository - GetByID : %w", err))
		return currency, err
	}

	return currency, nil
}

func (cr *CurrencyRepository) Update(id string, currency model.Currency) error {

	var sqlStatement = `
	UPDATE currencies 
	SET code= $2, description= $3
	WHERE id = $1
	`
	_, err := cr.DB.Exec(sqlStatement, id, currency.Code, currency.Description)
	if err != nil {
		log.Error(fmt.Errorf("error CurrencyRepository - Update : %w", err))
		return err
	}
	return nil
}

func (cr *CurrencyRepository) Delete(id string) error {

	var sqlStatement = `
	DELETE FROM currencies
	WHERE id = $1
	`

	_, err := cr.DB.Exec(sqlStatement, id)
	if err != nil {
		log.Error(fmt.Errorf("error CurrencyRepository - Delete : %w", err))
		return err
	}

	return nil
}
