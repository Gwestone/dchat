package DB

import "database/sql"

func MigrateUsers(db *sql.DB) {
	db.Exec("CREATE TABLE \"users\" (\n                         \"Id\"\tINTEGER NOT NULL UNIQUE,\n                         \"UserId\"\tTEXT NOT NULL UNIQUE,\n                         \"username\"\tTEXT NOT NULL UNIQUE,\n                         \"password\"\tTEXT NOT NULL,\n                         PRIMARY KEY(\"Id\" AUTOINCREMENT)\n);")
}
