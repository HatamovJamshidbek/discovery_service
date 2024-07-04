package postgres

import (
	"database/sql"
	"discovery_servcie/config"
	"fmt"
)

func ConnectionDb(conf *config.Config) (*sql.DB, error) {
	conDB := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable ", "localhost", 5433, "postgres", "discovery", "1111")
	db, err := sql.Open("postgres", conDB)
	if err != nil {
		return nil, err
	}
	return db, nil
}
