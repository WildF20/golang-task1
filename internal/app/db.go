package app

import (
	"database/sql"

	"golang-task1/config"
	"golang-task1/database"
)

func newDBConnection() (*sql.DB, error) {
	cfg, _ := config.LoadConfig()

	DBConn := "postgresql://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName

	db, err := database.InitDB(DBConn)
	if err != nil {
		return nil, err
	}
	// defer db.Close()

	return db, nil
}