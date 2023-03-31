package controllers

import (
	"Gawean/database"
	. "Gawean/helper"
	"Gawean/structs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetByDate(c *gin.Context) {
	var results []structs.Gawean
	x := c.Query("d")
	y := c.Query("m")
	z := c.Query("y")
	a := fmt.Sprintf("%s-%s-%s", z, y, x)
	sql := `SELECT *FROM gawean WHERE NextAppointment = $1`
	rows, err := database.DbConnection.Query(sql, a)
	CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		var b = structs.Gawean{}

		err = rows.Scan(&b.ID, &b.ClientID, &b.ClientName, &b.Description, &b.CreateDate, &b.UpdateDate, &b.CategoryID, &b.NextAppointment)
		CheckErr(err)
		results = append(results, b)

	}
	result := gin.H{
		"result": results,
	}

	c.JSON(http.StatusOK, result)

}
