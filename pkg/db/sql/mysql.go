package sql

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type MySqlDb struct {
	*sqlx.DB
}

func NewMysqlDB(config DBConfig) (*sqlx.DB, error) {

	mySqlDbInstance, err := openMysqlDBConnection(config)
	if err != nil {
		return nil, err
	}

	return mySqlDbInstance, nil
}

func openMysqlDBConnection(config DBConfig) (*sqlx.DB, error) {
	var err error

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.User, config.Pass, config.Host, MYSQLPORT, config.Name)
	mySqlDbInstance, err := sqlx.Open(MYSQL, connectionString)
	if err != nil {
		log.Println(err)
		return mySqlDbInstance, err
	}

	return mySqlDbInstance, nil
}
