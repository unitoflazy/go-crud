package main

import (
	"crud/controller"
	"crud/database"

	"github.com/gin-gonic/gin"
)

func SetUp() {
	database.ConnectDatabase()
}

func main() {

	SetUp()
	router := gin.Default()
	router.GET("/api/employees", controller.GetAllEmployees)
	router.GET("/api/employees/:id", controller.GetEmployee)
	router.DELETE("/api/employees/:id", controller.DeleteEmployee)
	router.PATCH("/api/employees/:id", controller.EditEmployee)
	router.POST("/api/employees", controller.CreateEmployee)

	router.Run(":8080")
}
