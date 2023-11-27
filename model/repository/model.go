package repository

import (
	"database/sql"
	"fmt"
	"log"
	"mycloud/config"
)

var Db *sql.DB

func init() {
	// configuration file
	cfg, _ := config.LoadConfigForYaml()

	var err error
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "127.0.0.1", 4000, "root", "secret", "mycloud")
	Db, err = sql.Open(cfg.Db.DbDriver, dataSourceName)
	if err != nil {
		log.Fatal("Invalid DB config:", err)
	}
	if err = Db.Ping(); err != nil {
		log.Fatal("DB unreachable:", err)
	}
}
