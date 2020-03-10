package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/isidroevc/monserver/configuration"
)

var database *sql.DB

func GetConnection() *sql.DB {
	if database == nil {
		config, readConfigError := configuration.GetConfiguration()
		if readConfigError != nil {
			panic(readConfigError)
		}
		db, err := sql.Open("mysql", config.MysqlConnectionString)
		database = db
		if err != nil {
			panic(err.Error())
		}
	}
	return database
}
