package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlDatabase(configuration Config) *sql.DB {
	connection := configuration.Get("DB_CONNECTION")
	host := configuration.Get("DB_HOST")
	username := configuration.Get("DB_USERNAME")
	password := configuration.Get("DB_PASSWORD")
	port := configuration.Get("DB_PORT")
	database := configuration.Get("DB_DATABASE")

	host_db := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
	connection_db := fmt.Sprintf("%s", connection)

	db, err := sql.Open(connection_db, host_db)
	//db, err := sql.Open("mysql", "root:@tcp(localhost:3306/golang_database)")

	if err != nil {
		panic(err)
	}
	return db
}
