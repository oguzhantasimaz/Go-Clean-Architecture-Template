package bootstrap

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

func NewMySQLDatabase(env *Env) *sqlx.DB {
	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass
	dbName := env.DBName

	// example connection string: "test:test@(localhost:3306)/test"

	db, err := sqlx.Connect("mysql", dbUser+":"+dbPass+"@("+dbHost+":"+dbPort+")/"+dbName+"?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CloseMySqlConnection(client *sqlx.DB) {
	if client == nil {
		return
	}

	err := client.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Connection to MySQL closed.")
}
