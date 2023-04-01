package middleware

import (
	"Gawean/database"
	. "Gawean/helper"
	"Gawean/structs"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticated(db *sql.DB, c *gin.Context) (auth bool) {
	auth = false
	un, pw, ok := c.Request.BasicAuth()

	sql := `SELECT *FROM Users`

	rows, err := db.Query(sql)
	CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		var x = structs.User{}

		err = rows.Scan(&x.Username, &x.Password)
		CheckErr(err)
		if ok && un == x.Username && pw == x.Password {
			auth = true
			c.Request.SetBasicAuth(un, pw)
			return
		}

	}
	return

}

func BasicAuth(ctx *gin.Context) {
	if Authenticated(database.DbConnection, ctx) {
		log.Println("Authenticated")
	} else {
		log.Println("Fail Authentication")
		ctx.JSON(http.StatusForbidden, gin.H{
			"Warning": "Not authorized",
		})
		ctx.Abort()
		return
	}
}
