package sql

import (
	"github.com/jmoiron/sqlx"
)

type DBConfig struct {
	Driver string
	Host   string
	User   string
	Pass   string
	Name   string
}

func NewDB(dbConfig *DBConfig) (*sqlx.DB, error) {

	switch dbConfig.Driver {
	case MYSQL:
		mySqlDbInstance, err := NewMysqlDB(*dbConfig)
		return mySqlDbInstance, err

	default:
		dbConErr := NewDBConnError(UNSUPPORTED_DB_DRIVER)
		return nil, dbConErr
	}
}
