package app

import (
	"database/sql"

	"golang-task1/config"
	"golang-task1/database"
)

func newDBConnection() (*sql.DB, error) {
	cfg, _ := config.LoadConfig()

	db, err := database.InitDB(cfg.DBConn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return db, nil
}