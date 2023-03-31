package controllers

import (
	"Gawean/database"
	"Gawean/repository"
	"Gawean/structs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllClients(c *gin.Context) {
	var result gin.H

	err, x := repository.GetAllClients(database.DbConnection)

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
func GetClient(c *gin.Context) {
	var result gin.H

	err, x := repository.GetClient(database.DbConnection, c)

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
func InsertClient(c *gin.Context) {
	var x structs.Client

	err := c.ShouldBindJSON(&x)
	if err != nil {
		panic(err)
	}

	err = repository.InsertClient(database.DbConnection, x)
	if err != nil {
		panic(err)
	}
	x.CreateDate = time.Time.String(time.Now())
	x.UpdateDate = time.Time.String(time.Now())

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Client",
		"Data":   x,
	})
}

func UpdateClient(c *gin.Context) {
	var x structs.Client
	id := c.Param("id")

	err := c.ShouldBindJSON(&x)
	if err != nil {
		panic(err)
	}
	x.ID = id
	err = repository.UpdateClient(database.DbConnection, x)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Client",
	})
}

func DeleteClient(c *gin.Context) {
	var x structs.Client
	id := c.Param("id")

	x.ID = id

	err := repository.DeleteClient(database.DbConnection, x)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Client",
	})
}

func DeleteClientSpecific(c *gin.Context) {
	var x structs.Client
	id := c.Param("id")
	name := c.Param("name")

	x.ID = id
	x.Name = name

	err := repository.DeleteClientSpecific(database.DbConnection, x)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Client",
	})
}
