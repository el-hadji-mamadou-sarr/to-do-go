package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"to-do-go/pkg/models"

	"github.com/gin-gonic/gin"
)

var tasks []models.Task

const taskFile = "tasks.json"

// Load tasks from file
func LoadTasks() error {
	file, err := os.Open(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			tasks = []models.Task{} // If file doesn't exist, start with empty tasks
			return nil
		}
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &tasks)
}

// Save tasks to file
func SaveTasks() error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(taskFile, data, 0644)
}

// Get all tasks
func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

// Create a new task
func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	if err := SaveTasks(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save task"})
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

// Delete a task
func DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			if err := SaveTasks(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save tasks"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// Update an existing task
func UpdateTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var taskIndex = -1
	for i, task := range tasks {
		if task.ID == id {
			taskIndex = i
			break
		}
	}

	if taskIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks[taskIndex].Title = updatedTask.Title
	tasks[taskIndex].Status = updatedTask.Status

	if err := SaveTasks(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks[taskIndex])
}
