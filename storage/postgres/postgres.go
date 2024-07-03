package postgres

import (
	"database/sql"
	"discovery_servcie/config"
	"fmt"
)

func ConnectionDb(conf *config.Config) (*sql.DB, error) {
	conDB := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable ", conf.PostgresHost, conf.PostgresPort, conf.PostgresUser, conf.PostgresDatabase, conf.PostgresPassword)
	db, err := sql.Open("postgres", conDB)
	if err != nil {
		return nil, err
	}
	return db, nil
}
