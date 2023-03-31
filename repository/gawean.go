package repository

import (
	. "Gawean/helper"
	"Gawean/structs"
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllGawean(db *sql.DB) (err error, results []structs.Gawean) {
	sql := `SELECT *FROM gawean`
	rows, err := db.Query(sql)
	CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var b = structs.Gawean{}

		err = rows.Scan(&b.ID, &b.ClientID, &b.ClientName, &b.Description, &b.CreateDate, &b.UpdateDate, &b.CategoryID, &b.NextAppointment)
		CheckErr(err)
		results = append(results, b)

	}
	return

}

func GetGaweanByClient(db *sql.DB, c *gin.Context) (err error, results []structs.Gawean) {
	sql := `SELECT *FROM gawean WHERE Clientid = $1`
	rows, err := db.Query(sql, c.Param("id"))
	CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var b = structs.Gawean{}

		err = rows.Scan(&b.ID, &b.ClientID, &b.ClientName, &b.Description, &b.CreateDate, &b.UpdateDate, &b.CategoryID, &b.NextAppointment)
		CheckErr(err)
		results = append(results, b)

	}
	return

}

func GetGaweanByCategory(db *sql.DB, c *gin.Context) (err error, results []structs.Gawean) {
	sql := `SELECT *FROM gawean WHERE CategoryID = $1`
	rows, err := db.Query(sql, c.Param("id"))
	CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var b = structs.Gawean{}

		err = rows.Scan(&b.ID, &b.ClientID, &b.ClientName, &b.Description, &b.CreateDate, &b.UpdateDate, &b.CategoryID, &b.NextAppointment)
		CheckErr(err)
		results = append(results, b)

	}
	return

}

func InsertGawean(db *sql.DB, g structs.Gawean) (er error) {
	sql := `INSERT INTO gawean (ID, ClientID, ClientName, Descr, CreateDate, UpdateDate, CategoryID, NextAppointment) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	errs := db.QueryRow(sql, g.ID, g.ClientID, g.ClientName, g.Description, time.Now(), time.Now(), g.CategoryID, g.NextAppointment)

	return errs.Err()
}

func UpdateGawean(db *sql.DB, g structs.Gawean, c *gin.Context) (er error) {
	sql := `UPDATE gawean SET ID = $1, ClientID = $2, ClientName = $3, Descr = $4, UpdateDate =$5, CategoryID =$6, NextAppointment = $7 WHERE ID=$8`

	errs := db.QueryRow(sql, g.ID, g.ClientID, g.ClientName, g.Description, time.Now(), g.CategoryID, g.NextAppointment, c.Param("id"))

	return errs.Err()
}

func DeleteGawean(db *sql.DB, g structs.Gawean) (er error) {
	sql := `DELETE FROM gawean WHERE ID = $1`

	errs := db.QueryRow(sql, g.ID)

	return errs.Err()
}
