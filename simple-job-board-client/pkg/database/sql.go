package database

import (
	"database/sql"
	"fmt"
	"time"
)

// Temporary hardcoded postgresql
func GetDbConn(databaseUser string, databasePassword string, databaseHost string, databasePort string, databaseName string, sslMode string) (*sql.DB, error) {
	databaseURL := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%s",
		databaseUser,
		databasePassword,
		databaseHost,
		databasePort,
		databaseName,
		sslMode,
	)
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(5 * time.Minute)
	return db, nil
}

func CloseDbConn(conn *sql.DB) {
	conn.Close()
}
