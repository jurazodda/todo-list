package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

var tasks []Task
var nextID int = 1

func main() {
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

	task := Task{
		ID:     nextID,
		Title:  input.Title,
		IsDone: false,
	}
	nextID++

	tasks = append(tasks, task)

	c.JSON(http.StatusCreated, gin.H{"task": task})
}

func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
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

	for i := range tasks {
		if tasks[i].ID == id {
			c.JSON(http.StatusOK, tasks[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
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

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Title = input.Title
			c.JSON(http.StatusOK, gin.H{"message": "task updated"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
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

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].IsDone = true
			c.JSON(http.StatusOK, gin.H{"message": "task completed", "tasks": tasks[i]})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
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

	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
