package app

import (
	"golang-task1/config"
	"golang-task1/database"
)

func newDBConnection() error {
	cfg, _ := config.LoadConfig()

	db, err := database.InitDB(cfg.DBConn)
	if err != nil {
		return err
	}
	defer db.Close()

	return nil
}