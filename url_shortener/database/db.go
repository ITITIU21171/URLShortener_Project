package database

import (
	"database/sql"
	"log"

	"github.com/ITITIU21171/url-shortener/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	cfg := config.LoadConfig()

	var err error
	DB, err = sql.Open("postgres", cfg.DB_URL)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Database is not reachable:", err)
	}

	log.Println("Connected to the database successfully!")
}
