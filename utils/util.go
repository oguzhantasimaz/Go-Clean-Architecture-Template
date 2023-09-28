package utils

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func JSON(w http.ResponseWriter, code int, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.Encode(obj)
}

func MigrateDB(db *sqlx.DB) {
	db.MustExec(`
		CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		google_id VARCHAR(255) DEFAULT '',
		profile_picture VARCHAR(255) DEFAULT '',
		name VARCHAR(255) DEFAULT '',
		password VARCHAR(255) DEFAULT '',
		email VARCHAR(255) NOT NULL UNIQUE,
		phone VARCHAR(255) DEFAULT '',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP 
		);
	`)
}

func SetCookie(w http.ResponseWriter, name string, value string) {
	cookie := http.Cookie{
		Name:  name,
		Value: value,
		Path:  "/",
	}
	http.SetCookie(w, &cookie)
}
