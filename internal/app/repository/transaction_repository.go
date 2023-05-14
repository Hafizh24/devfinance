package repository

import (
	"fmt"

	"github.com/hafizh24/devfinance/internal/app/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type TransactionRepository struct {
	DB *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (tr *TransactionRepository) Create(Transaction model.Transaction) error {
	var sqlStatement = `
		INSERT INTO transactions (type, note, amount, category_id, currency_id, user_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		`

	_, err := tr.DB.Exec(sqlStatement, Transaction.Type, Transaction.Note, Transaction.Amount, Transaction.CategoryID, Transaction.CurrencyID, Transaction.UserID)
	if err != nil {
		log.Error(fmt.Errorf("error Transaction Repository - Create : %w", err))
		return err
	}
	return nil
}

func (tr *TransactionRepository) Browse(UserID int) ([]model.Transaction, error) {

	var (
		transactions []model.Transaction
		sqlStatement = `
	SELECT transactions.id,transactions.type,transactions.note,transactions.amount,transactions.created_at,categories.name,currencies.code || ' ' || transactions.amount AS total_amount
	FROM transactions
	INNER JOIN categories ON transactions.category_id = categories.id
	INNER JOIN currencies ON transactions.currency_id = currencies.id
	WHERE transactions.user_id = $1
	ORDER BY transactions.created_at DESC
	`
	)
	rows, err := tr.DB.Queryx(sqlStatement, UserID)
	if err != nil {
		log.Error(fmt.Errorf("error Transaction Repository - Browse : %w", err))
		return transactions, err
	}

	for rows.Next() {
		var Transaction model.Transaction
		// nolint:errcheck
		rows.StructScan(&Transaction)
		transactions = append(transactions, Transaction)
	}
	return transactions, nil
}

func (tr *TransactionRepository) GetByID(id string) (model.Transaction, error) {
	var Transaction model.Transaction
	var sqlStatement = `
	SELECT id,type,note,amount,category_id,currency_id
	FROM transactions
	WHERE id = $1
	`
	err := tr.DB.QueryRowx(sqlStatement, id).StructScan(&Transaction)
	if err != nil {
		log.Error(fmt.Errorf("error Transaction Repository - GetByID : %w", err))
		return Transaction, err
	}

	return Transaction, nil
}
func (tr *TransactionRepository) GetByType(Type string, UserID int) ([]model.Transaction, error) {
	var transactions []model.Transaction

	var sqlStatement = `
	SELECT transactions.id,transactions.type,transactions.note,transactions.amount,transactions.created_at,categories.name,currencies.code || ' ' || transactions.amount AS total_amount
	FROM transactions
	INNER JOIN categories ON transactions.category_id = categories.id
	INNER JOIN currencies ON transactions.currency_id = currencies.id
	WHERE transactions.type = $1 AND transactions.user_id = $2
	`

	rows, err := tr.DB.Queryx(sqlStatement, Type, UserID)

	if err != nil {
		log.Error(fmt.Errorf("error Transaction Repository - GetByType : %w", err))
		return transactions, err
	}

	for rows.Next() {
		var Transaction model.Transaction
		// nolint:errcheck
		rows.StructScan(&Transaction)
		transactions = append(transactions, Transaction)
	}

	return transactions, nil
}

func (tr *TransactionRepository) Update(id string, Transaction model.Transaction) error {

	var sqlStatement = `
	UPDATE transactions 
	SET type = $2, note = $3, amount = $4, category_id = $5, currency_id = $6
	WHERE id = $1
	`
	_, err := tr.DB.Exec(sqlStatement, id, Transaction.Type, Transaction.Note, Transaction.Amount, Transaction.CategoryID, Transaction.CurrencyID)
	if err != nil {
		log.Error(fmt.Errorf("error Transaction Repository - Update : %w", err))
		return err
	}
	return nil
}

func (tr *TransactionRepository) Delete(id string) error {

	var sqlStatement = `
	DELETE FROM transactions
	WHERE id = $1
	`

	_, err := tr.DB.Exec(sqlStatement, id)
	if err != nil {
		log.Error(fmt.Errorf("error Transaction Repository - Delete : %w", err))
		return err
	}

	return nil
}
