package repository

import (
	. "Gawean/helper"
	"Gawean/structs"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func GetAllAddresses(db *sql.DB) (err error, results []structs.Address) {
	sql := `SELECT *FROM address`

	rows, err := db.Query(sql)
	CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var b = structs.Address{}

		err = rows.Scan(&b.ID, &b.Province, &b.City, &b.Description)
		CheckErr(err)
		results = append(results, b)

	}
	return

}

func GetAddress(db *sql.DB, c *gin.Context) (err error, results []structs.Address) {
	sql := `SELECT *FROM address WHERE ID = $1`
	rows, err := db.Query(sql, c.Param("id"))
	CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var b = structs.Address{}

		err = rows.Scan(&b.ID, &b.Province, &b.City, &b.Description)
		CheckErr(err)
		results = append(results, b)

	}
	return

}

func InsertAddress(db *sql.DB, ad structs.Address) (er error) {
	sql := `INSERT INTO address (ID, Province, City, Descr) VALUES ($1, $2, $3, $4)`

	errs := db.QueryRow(sql, ad.ID, ad.Province, ad.City, ad.Description)

	return errs.Err()
}

func UpdateAddress(db *sql.DB, ad structs.Address, c *gin.Context) (er error) {
	sql := `UPDATE address SET ID = $1, Province = $2, City = $3, Descr = $4 WHERE ID =$5`

	errs := db.QueryRow(sql, ad.ID, ad.Province, ad.City, ad.Description, c.Param("id"))

	return errs.Err()
}

func DeleteAddress(db *sql.DB, ad structs.Address) (er error) {
	sql := `DELETE FROM address WHERE id = $1`

	errs := db.QueryRow(sql, ad.ID)

	return errs.Err()
}
