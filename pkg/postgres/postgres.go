package postgres

import (
	"fmt"
	"go-clean-arch/config"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	_ "github.com/jackc/pgx/stdlib"
)

// InitPostgres DB Instance
func InitPostgres(conf *config.Config) (*sqlx.DB, error) {

	source := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Postgres.Hostname, conf.Postgres.Port, conf.Postgres.Username, conf.Postgres.Password, conf.Postgres.DBName)

	db, err := connectDB(conf.Postgres.DriverName, source)
	if err != nil {
		return nil, err
	}

	return db, nil

}

func connectDB(driverName string, dataSourceName string) (*sqlx.DB, error) {

	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, errors.New("database postgres connect")
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)
	db.SetConnMaxIdleTime(time.Minute * 5)
	db.SetConnMaxLifetime(time.Minute * 5)

	if err = db.Ping(); err != nil {
		return nil, errors.New("database postgres ping")
	}

	return db, nil

}
