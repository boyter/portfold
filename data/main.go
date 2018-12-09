package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() (*sql.DB, error) {
	// Connect to the database
	// parseTime is required to ensure that DATETIME to time.Time conversions work in SQL mapper
	db, err := openDb("root:@tcp(localhost:3306)/portfold?autocommit=true&parseTime=true")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func openDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(5)

	return db, nil
}
