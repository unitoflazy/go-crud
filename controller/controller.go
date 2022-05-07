package controller

import (
	"crud/database"
	"crud/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllEmployees(c *gin.Context) {
	db := database.Pool

	var employees []model.Employee

	result := db.Find(&employees)

	if result.Error != nil {
		sendFailed(c, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"data":    employees,
		"message": "",
	})
}

func GetEmployee(c *gin.Context) {
	db := database.Pool
	id := c.Param("id")
	var employee model.Employee

	result := db.Where("id = ?", id).Find(&employee)

	if result.Error != nil {
		sendFailed(c, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		sendFailed(c, fmt.Sprintf("Employee ID %v not found\n", id), http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"data":    employee,
		"message": "",
	})
}

func CreateEmployee(c *gin.Context) {
	db := database.Pool
	var employee model.Employee

	if err := c.ShouldBind(&employee); err != nil {
		sendFailed(c, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.Select("name", "age").Create(&employee)

	if result.Error != nil {
		sendFailed(c, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"data":    model.Employee{ID: employee.ID, Name: employee.Name, Age: employee.Age},
		"message": "Created",
	})
}

func DeleteEmployee(c *gin.Context) {
	db := database.Pool
	id := c.Param("id")

	result := db.Where("id = ?", id).Delete(&model.Employee{})

	if result.Error != nil {
		sendFailed(c, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		sendFailed(c, fmt.Sprintf("Employee ID %v not found\n", id), http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"data":    "",
		"message": "Deleted",
	})
}

func EditEmployee(c *gin.Context) {
	db := database.Pool
	id, _ := strconv.Atoi(c.Param("id"))

	var employee model.Employee

	if err := c.ShouldBind(&employee); err != nil {
		sendFailed(c, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.Table("employees").Where("id = ?", id).Updates(map[string]interface{}{
		"name": employee.Name,
		"age":  employee.Age,
	})

	if result.Error != nil {
		sendFailed(c, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		sendFailed(c, fmt.Sprintf("Employee ID %v not found\n", id), http.StatusNotFound)
		return
	}

	employee.ID = id

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"data":    employee,
		"message": "Updated",
	})
}

func sendFailed(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, gin.H{
		"status":  "failed",
		"message": message,
	})
}
