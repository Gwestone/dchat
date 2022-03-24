package DB

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() *sql.DB {

	//db, err := sql.Open("postgres", "postgresql://dchat:726nk6WwJf3J6hFuSymq3nexwgHdYPSA@localhost:5432/dchat")
	db, err := sql.Open("sqlite3", "./dchat.db")
	if err != nil {
		panic(err.Error())
	}
	return db

}
