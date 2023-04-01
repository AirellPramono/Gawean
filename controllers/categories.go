package controllers

import (
	"Gawean/database"
	. "Gawean/helper"
	"Gawean/repository"
	"Gawean/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	var result gin.H

	err, x := repository.GetAllCategories(database.DbConnection)

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

func InsertCategory(c *gin.Context) {
	var x structs.Category

	err := c.ShouldBindJSON(&x)
	CheckErr(err)

	err = repository.InsertCategory(database.DbConnection, x)
	CheckErr(err)

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Category",
		"Data":   x,
	})
}

func UpdateCategory(c *gin.Context) {
	var x structs.Category
	err := c.ShouldBindJSON(&x)
	CheckErr(err)
	err = repository.UpdateCategory(database.DbConnection, x, c)

	CheckErr(err)

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Category",
	})
}

func DeleteCategory(c *gin.Context) {
	var x structs.Category
	id, err := strconv.Atoi(c.Param("id"))

	x.ID = id

	err = repository.DeleteCategory(database.DbConnection, x)

	CheckErr(err)

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Category",
	})
}
