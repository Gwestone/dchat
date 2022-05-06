package DB

import (
	"database/sql"
	"fmt"
	"io/ioutil"
)

func MigrateUsers(db *sql.DB) {

	b, err := ioutil.ReadFile("../SQL/dchat_localhost-2022_05_06_21_14_56-dump.sql") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	queryStr := string(b) // convert content to a 'string'
	fmt.Print(queryStr)
	_, err = db.Exec(queryStr)
	if err != nil {
		fmt.Print("failed to migrate db: " + err.Error())
	}
}
