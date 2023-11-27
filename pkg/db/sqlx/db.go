package sqlx

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type DbConfig struct {
	Driver, Host, User, Pass, Name string
}

func NewSqlDB(dbConf *DbConfig) (*sqlx.DB, error) {

	switch dbConf.Driver {
	case MYSQL:
		mysqlDbInstance, err := NewMysqlDB(dbConf.Host, dbConf.User, dbConf.Pass, dbConf.Name)
		return mysqlDbInstance, err
	default:
		return nil, errors.New("database driver not supported")
	}

}
