package handlers

import (
	"net/http"
	"to-do-go/pkg/models"

	"github.com/gin-gonic/gin"
)
import "strconv"

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
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de tâche invalide"})
		return
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func UpdateTask(c *gin.Context) {
	// Récupérer l'ID depuis les paramètres d'URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de tâche invalide"})
		return
	}

	// Rechercher la tâche correspondante
	var taskIndex = -1
	for i, task := range tasks {
		if task.ID == id {
			taskIndex = i
			break
		}
	}
	if taskIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tâche introuvable"})
		return
	}

	// Analyser le corps de la requête
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mettre à jour les champs modifiables (sauf l'ID)
	tasks[taskIndex].Title = updatedTask.Title
	tasks[taskIndex].Status = updatedTask.Status
	// Ajouter ici d'autres champs à mettre à jour si nécessaire

	c.JSON(http.StatusOK, tasks[taskIndex])
}