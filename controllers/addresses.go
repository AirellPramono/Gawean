package controllers

import (
	"Gawean/database"
	"Gawean/repository"
	"Gawean/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAddresses(c *gin.Context) {
	var result gin.H

	err, x := repository.GetAllAddresses(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": x,
		}
	}
	c.JSON(http.StatusOK, result)
}

func GetAddress(c *gin.Context) {
	var result gin.H

	err, x := repository.GetAddress(database.DbConnection, c)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": x,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertAddress(c *gin.Context) {
	var x structs.Address

	err := c.ShouldBindJSON(&x)
	if err != nil {
		panic(err)
	}

	err = repository.InsertAddress(database.DbConnection, x)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Address",
		"Data":   x,
	})
}

func UpdateAddress(c *gin.Context) {
	var x structs.Address
	err := c.ShouldBindJSON(&x)
	if err != nil {
		panic(err)
	}
	err = repository.UpdateAddress(database.DbConnection, x, c)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Address",
	})
}

func DeleteAddress(c *gin.Context) {
	var x structs.Address
	id := c.Param("id")

	x.ID = id

	err := repository.DeleteAddress(database.DbConnection, x)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Address",
	})
}
