package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Task struct {
	gorm.Model
	Title  string `json:"title" gorm:"not null"`
	IsDone bool   `json:"is_done"`
}

func initDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=todo-list port=5433 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database" + err.Error())
	}

	db.AutoMigrate(&Task{})
}

func main() {
	initDB()
	router := gin.Default()

	router.POST("/tasks", createTask)
	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTaskByID)
	router.PATCH("/tasks/:id", updateTask)
	router.PATCH("/tasks/:id/complete", completeTask)
	router.DELETE("/tasks/:id", deleteTask)

	router.Run(":8080")
}

func createTask(c *gin.Context) {
	input := struct {
		Title string `json:"title"`
	}{}

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	task := Task{Title: input.Title}
	db.Create(&task)

	c.JSON(http.StatusCreated, gin.H{"task": task})
}

func getTasks(c *gin.Context) {
	var tasks []Task
	db.Find(&tasks)
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func getTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task id required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed convert to int"})
		return
	}

	var task Task
	if err := db.WithContext(c).First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func updateTask(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task id required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed convert to int"})
		return
	}

	input := struct {
		Title string `json:"title"`
	}{}

	err = c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var task Task
	if err := db.WithContext(c).First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}

	task.Title = input.Title
	db.Save(&task)

	c.JSON(http.StatusOK, task)
}

func completeTask(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task id required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed convert to int"})
		return
	}

	var task Task
	if err := db.WithContext(c).First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}

	task.IsDone = true
	db.Save(&task)

	c.JSON(http.StatusOK, gin.H{"message": "Task completed", "task": task})
}

func deleteTask(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task id required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed convert to int"})
		return
	}

	res := db.Delete(&Task{}, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
