package repository

import (
	"fmt"

	"github.com/hafizh24/devfinance/internal/app/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type CategoryRepository struct {
	DB *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (cr *CategoryRepository) Create(category model.Category) error {
	var sqlStatement = `
		INSERT INTO categories (name, description)
		VALUES ($1, $2)
		`

	_, err := cr.DB.Exec(sqlStatement, category.Name, category.Description)
	if err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - Create : %w", err))
		return err
	}
	return nil
}

func (cr *CategoryRepository) Browse() ([]model.Category, error) {

	var categories []model.Category
	var sqlStatement = `
		SELECT id,name,description
		FROM categories
	`
	rows, err := cr.DB.Queryx(sqlStatement)
	if err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - Browse : %w", err))
		return categories, err
	}

	for rows.Next() {
		var category model.Category
		// nolint:errcheck
		rows.StructScan(&category)
		categories = append(categories, category)
	}
	return categories, nil
}

func (cr *CategoryRepository) GetByID(id string) (model.Category, error) {
	var category model.Category
	var sqlStatement = `
	SELECT id,name,description
	FROM categories
	WHERE id = $1
	`
	err := cr.DB.QueryRowx(sqlStatement, id).StructScan(&category)
	if err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - GetByID : %w", err))
		return category, err
	}

	return category, nil
}

func (cr *CategoryRepository) Update(id string, category model.Category) error {

	var sqlStatement = `
	UPDATE categories 
	SET name= $2, description= $3
	WHERE id = $1
	`
	_, err := cr.DB.Exec(sqlStatement, id, category.Name, category.Description)
	if err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - Update : %w", err))
		return err
	}
	return nil
}

func (cr *CategoryRepository) Delete(id string) (model.Category, error) {
	var category model.Category
	var sqlStatement = `
	DELETE FROM categories
	WHERE id = $1
	`

	_, err := cr.DB.Exec(sqlStatement, id)
	if err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - Delete : %w", err))
		return category, err
	}

	return category, nil
}
