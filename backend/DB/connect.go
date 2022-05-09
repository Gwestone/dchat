package DB

import (
	"database/sql"
	_ "github.com/lib/pq"
)

//_ "github.com/mattn/go-sqlite3"

func ConnectDB(dbUrl string) (*sql.DB, error) {

	db, err := sql.Open("postgres", dbUrl)
	//db, err := sql.Open("sqlite3", "./dchat.db")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil

}
