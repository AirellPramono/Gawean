package controllers

import (
	"Gawean/database"
	"Gawean/repository"
	"Gawean/structs"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllGawean(c *gin.Context) {
	var result gin.H

	err, x := repository.GetAllGawean(database.DbConnection)

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

func InsertGawean(c *gin.Context) {
	var x structs.Gawean

	err := c.ShouldBindJSON(&x)
	if err != nil {
		panic(err)
	}

	err = repository.InsertGawean(database.DbConnection, x)
	if err != nil {
		panic(err)
	}

	x.CreateDate = time.Time.String(time.Now())
	x.UpdateDate = time.Time.String(time.Now())

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Gawean",
		"Data":   x,
	})
}

func UpdateGawean(c *gin.Context) {
	var x structs.Gawean
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&x)
	if err != nil {
		panic(err)
	}
	x.ID = id
	err = repository.UpdateGawean(database.DbConnection, x)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Gawean",
	})
}

func DeleteGawean(c *gin.Context) {
	var x structs.Gawean
	id, err := strconv.Atoi(c.Param("id"))

	x.ID = id

	err = repository.DeleteGawean(database.DbConnection, x)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Gawean",
	})
}
