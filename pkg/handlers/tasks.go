package handlers

import (
	"net/http"
	"to-do-go/pkg/models"

	"github.com/gin-gonic/gin"
)

// In-memory task list (temporary storage)
var tasks = []models.Task{}

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, newTask)
}

func DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
