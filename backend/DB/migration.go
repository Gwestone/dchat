package DB

import (
	"database/sql"
	"fmt"
	"io/ioutil"
)

func MigrateUsers(db *sql.DB) {

	rawSql, err := ioutil.ReadFile("./End_migrate.sql") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	queryStr := string(rawSql) // convert content to a 'string'
	//fmt.Print(queryStr)
	_, err = db.Exec(queryStr)
	if err != nil {
		fmt.Print("failed to migrate db: " + err.Error())
	}
}
