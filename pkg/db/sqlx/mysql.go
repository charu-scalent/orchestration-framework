package sqlx

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySqlDb struct {
	*sqlx.DB
}

func NewMysqlDB(dbHost string, dbUser string, dbPass string, dbName string) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, MYSQLPORT, dbName)
	mysqlDbInstance, err := sqlx.Connect(MYSQL, connectionString)
	if err != nil {
		log.Println(err) // logger
		return nil, err
	}
	return mysqlDbInstance, nil
}
