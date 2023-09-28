package bootstrap

import (
	"github.com/jmoiron/sqlx"
)

type Application struct {
	Env   *Env
	MySql *sqlx.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.MySql = NewMySQLDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMySqlConnection(app.MySql)
}
