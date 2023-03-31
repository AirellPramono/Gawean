package repository

import (
	. "Gawean/helper"
	"Gawean/structs"
	"database/sql"
	"time"
)

func GetAllGawean(db *sql.DB) (err error, results []structs.Gawean) {
	sql := `SELECT *FROM gawean`
	rows, err := db.Query(sql)
	CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var b = structs.Gawean{}

		err = rows.Scan(&b.ID, &b.ClientID, &b.ClientName, &b.Description, &b.CreateDate, &b.UpdateDate)
		CheckErr(err)
		results = append(results, b)

	}
	return

}

func InsertGawean(db *sql.DB, g structs.Gawean) (er error) {
	sql := `INSERT INTO gawean (ID, ClientID, ClientName, Descr, CreateDate, UpdateDate) VALUES ($1, $2, $3, $4, $5, $6)`

	errs := db.QueryRow(sql, g.ID, g.ClientID, g.ClientName, g.Description, time.Now(), time.Now())

	return errs.Err()
}

func UpdateGawean(db *sql.DB, g structs.Gawean) (er error) {
	sql := `UPDATE gawean SET ID = $1, ClientID = $2, ClientName = $3, Descr = $4, UpdateDate =$5`

	errs := db.QueryRow(sql, g.ClientID, g.ClientName, g.Description, time.Now())

	return errs.Err()
}

func DeleteGawean(db *sql.DB, g structs.Gawean) (er error) {
	sql := `DELETE FROM gawean WHERE id = $1`

	errs := db.QueryRow(sql, g.ID)

	return errs.Err()
}
