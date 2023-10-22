package postgresql

import (
	"database/sql"
	"effective-test/pkg/logger"
)

func NewClient(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		logger.Errorf("sql.Open err %v", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		logger.Errorf("db.Ping err %v", err)
		return nil, err
	}

	return db, nil
}
