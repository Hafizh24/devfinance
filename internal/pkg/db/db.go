package db

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

func ConnectDB(DBDriver string, DBConn string) (*sqlx.DB, error) {

	db, err := sqlx.Connect(DBDriver, DBConn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Info("Database connection established.")

	return db, nil
}
