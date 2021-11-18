package controller

import (
	"log"
	"strconv"

	"github.com/emersonluiz/go-user/db"
	"github.com/emersonluiz/go-user/models"
	"github.com/emersonluiz/go-user/repository"
	"github.com/emersonluiz/go-user/service"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user *models.User
	connection := db.Connect()
	c.BindJSON(&user)
	rtn, err := repository.CreateUser(connection, user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot insert user on DB",
		})
		return
	}
	service.SendMessage(user)
	c.JSON(201, rtn)
}

func FindAllUser(c *gin.Context) {
	var user *models.User
	connection := db.Connect()
	c.BindJSON(&user)
	rtn, err := repository.FindAllUser(connection)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"error": "Cannot list users",
		})
		return
	}
	c.JSON(200, rtn)
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(404, gin.H{
			"error": "Cannot converd parameter to int",
		})
		return
	}
	connection := db.Connect()
	err = repository.DeleteUser(connection, id)
	if err != nil {
		log.Println(err)
		c.JSON(404, gin.H{
			"error": "User was not found",
		})
		return
	}
	c.Writer.WriteHeader(204)
}

func FindOneUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(404, gin.H{
			"error": "Cannot converd parameter to int",
		})
		return
	}
	connection := db.Connect()
	rtn, error := repository.FindOneUser(connection, id)
	if error != nil {
		log.Println(err)
		c.JSON(404, gin.H{
			"error": "User was not found",
		})
		return
	}

	c.JSON(200, rtn)
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(404, gin.H{
			"error": "Cannot converd parameter to int",
		})
		return
	}
	var user *models.User
	c.BindJSON(&user)
	connection := db.Connect()
	err = repository.UpdateUser(connection, id, user)
	if err != nil {
		log.Println(err)
		c.JSON(404, gin.H{
			"error": "User was not found",
		})
		return
	}
	c.Writer.WriteHeader(204)
}
