package repository

import (
	"Gawean/structs"
	"database/sql"
)

func RegisterUser(db *sql.DB, us structs.User) (er error) {
	sql := `INSERT INTO users (un, pw) VALUES ($1, $2)`

	errs := db.QueryRow(sql, us.Username, us.Password)

	return errs.Err()
}
