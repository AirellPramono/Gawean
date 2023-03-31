package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuth(ctx *gin.Context) {
	un, pw, auth := ctx.Request.BasicAuth()
	if auth && un == "Arif" && pw == "Bugaresa" {
		log.Println("Authenticated")
	} else if auth && un == "Saya" && pw == "Ganteng" {
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
