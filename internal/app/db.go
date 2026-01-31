package app

import (
	"database/sql"
	"log"
	"os"

	"golang-task1/config"
	"golang-task1/database"
)

func newDBConnection() (*sql.DB, error) {
	cfg, _ := config.LoadConfig()

	DBConn := "postgresql://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName
	log.Println("DB_CONN =", DBConn)
	log.Println("DB_HOST =", os.Getenv("DB_HOST"))
	log.Println("DB_PORT =", os.Getenv("DB_PORT"))


	db, err := database.InitDB(DBConn)
	if err != nil {
		return nil, err
	}
	// defer db.Close()

	return db, nil
}