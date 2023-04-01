package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresConnection struct {
}

func GetConnection() (*sql.DB, error) {
	return sql.Open("postgres", "user=postgres password=postgres dbname=food sslmode=disable")
}

func (c PostgresConnection) Query(query string, args []([]interface{})) error {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=food sslmode=disable")

	if err != nil {
		return err
	}

	defer db.Close()

	var result *sql.Row

	if len(args) > 1 {
		result = db.QueryRow(query, args[0]...)
	} else {
		result = db.QueryRow(query)
	}

	if result != nil {
		if err := result.Scan(args[1]...); err != nil {
			return err
		}

		return nil
	}

	return nil
}
