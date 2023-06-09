package repository

import (
	. "Gawean/helper"
	"Gawean/structs"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(db *sql.DB) (err error, results []structs.Category) {
	sql := `SELECT *FROM category`

	rows, err := db.Query(sql)
	CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var b = structs.Category{}

		err = rows.Scan(&b.ID, &b.Name)
		CheckErr(err)
		results = append(results, b)

	}
	return

}

func InsertCategory(db *sql.DB, ct structs.Category) (er error) {
	sql := `INSERT INTO category (ID, Fullname) VALUES ($1, $2)`

	errs := db.QueryRow(sql, ct.ID, ct.Name)

	return errs.Err()
}

func UpdateCategory(db *sql.DB, ct structs.Category, c *gin.Context) (er error) {
	sql := `UPDATE category SET ID = $1, Fullname = $2 WHERE ID = $3`

	errs := db.QueryRow(sql, ct.ID, ct.Name, c.Param("id"))

	return errs.Err()
}

func DeleteCategory(db *sql.DB, ct structs.Category) (er error) {
	sql := `DELETE FROM category WHERE id = $1`

	errs := db.QueryRow(sql, ct.ID)

	return errs.Err()
}
