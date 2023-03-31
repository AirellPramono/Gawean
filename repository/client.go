package repository

import (
	. "Gawean/helper"
	"Gawean/structs"
	"database/sql"
	"time"
)

func GetAllClients(db *sql.DB) (err error, results []structs.Client) {
	sql := `SELECT *FROM client`

	rows, err := db.Query(sql)
	CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var b = structs.Client{}

		err = rows.Scan(&b.ID, &b.Name, &b.Description, &b.CreateDate, &b.UpdateDate, &b.City, &b.Province, &b.AddressID)
		CheckErr(err)
		results = append(results, b)

	}
	return

}

func InsertClient(db *sql.DB, cl structs.Client) (er error) {
	sql := `INSERT INTO client (ID, FullName, Descr, CreateDate, UpdateDate, City, Province, AddressID) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	errs := db.QueryRow(sql, cl.ID, cl.Name, cl.Description, time.Now(), time.Now(), cl.City, cl.Province, cl.AddressID)

	return errs.Err()
}

func UpdateClient(db *sql.DB, cl structs.Client) (er error) {
	sql := `UPDATE client SET ID = $1, FullName = $2, Descr = $3, UpdateDate = $4, City = $5, Province = $6, AddressID = $7`

	errs := db.QueryRow(sql, cl.ID, cl.Name, cl.Description, time.Now(), cl.City, cl.Province, cl.AddressID)

	return errs.Err()
}

func DeleteClient(db *sql.DB, cl structs.Client) (er error) {
	sql := `DELETE FROM client WHERE id = $1`

	errs := db.QueryRow(sql, cl.ID)

	return errs.Err()
}
