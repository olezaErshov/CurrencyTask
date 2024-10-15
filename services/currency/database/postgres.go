package database

import (
	"CurrencyTask/services/currency/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"log"
)

func NewDB(cfg config.DBConfig) (*sqlx.DB, error) {
	conStr := fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := sqlx.Connect("postgres", conStr)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to database")
	err = applyMigrations(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}
func applyMigrations(db *sqlx.DB) error {
	if err := goose.Up(db.DB, "services/currency/database/migrations"); err != nil {
		log.Println("Error applying migrations")
		return err
	}
	log.Println("Applied migrations")
	return nil
}
