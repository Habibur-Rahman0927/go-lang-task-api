package main

import (
	"fmt"
	"go-api/config"
	"go-api/controllers"
	"go-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// SetupDB : initializing mysql database

func main() {
	fmt.Println("Hello World")
	// core fundamental pratices
	// gopractice.Pracitce()

	r := gin.Default()
	db := config.SetupDB()
	db.AutoMigrate(&models.Task{})
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/test-route", bodyFunc)

	r.GET("/tasks", controllers.FindTasks)
	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks/:id", controllers.FindTask)
	r.PATCH("/tasks/:id", controllers.UpdateTask)
	r.DELETE("tasks/:id", controllers.DeleteTask)
	r.Run()
}

func bodyFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello world",
	})
}
