package databases

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {

	dsn := "host=covid-survey.cju68sqkut9o.ap-southeast-1.rds.amazonaws.com port=5432 user=postgres password=password dbname=postgres"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
