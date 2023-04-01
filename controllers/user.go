package controllers

import (
	"Gawean/database"
	. "Gawean/helper"
	"Gawean/repository"
	"Gawean/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var x structs.User

	err := c.ShouldBindJSON(&x)
	CheckErr(err)

	err = repository.RegisterUser(database.DbConnection, x)
	CheckErr(err)

	c.JSON(http.StatusOK, gin.H{
		"result": "User Registered",
		"Data":   x,
	})
}
