package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
	"to-do-go/pkg/models"

	"github.com/gin-gonic/gin"
)

var (
	tasks []models.Task
	mu    sync.Mutex // Ensures thread safety if accessed concurrently
)

const taskFile = "tasks.json"

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

func SaveTasks() error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(taskFile, data, 0644)
}

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

// CreateTask - Creates a new task
func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	taskID := len(tasks) + 1
	createdTask := models.NewTask(taskID, newTask.Title)

	tasks = append(tasks, *createdTask)

	if err := SaveTasks(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save task"})
		return
	}

	c.JSON(http.StatusCreated, createdTask)
}

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

func ProcessTask(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{"message": "Le traitement de la tâche a commencé", "task_id": id})

	go func(taskID string) {
		fmt.Printf("[Tâche %s] Début du traitement...\n", taskID)

		// set the task status to pending
		for i, task := range tasks {
			if strconv.Itoa(task.ID) == taskID {
				tasks[i].Status = models.Pending
				break
			}
		}

		time.Sleep(1 * time.Second)
		fmt.Printf("[Tâche %s] Étape 1: Initialisation terminée.\n", taskID)

		time.Sleep(2 * time.Second)
		fmt.Printf("[Tâche %s] Étape 2: Traitement en cours...\n", taskID)

		time.Sleep(1 * time.Second)
		fmt.Printf("[Tâche %s] Étape 3: Finalisation...\n", taskID)

		// set the task status to completed
		for i, task := range tasks {
			if strconv.Itoa(task.ID) == taskID {
				tasks[i].Status = models.Completed
				break
			}
		}

		time.Sleep(1 * time.Second)
		fmt.Printf("[Tâche %s] Traitement terminé avec succès !\n", taskID)
	}(id)
}

func ProcessTasksParallel(c *gin.Context) {
	var taskIDs = []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup
	wg.Add(len(taskIDs))

	for _, id := range taskIDs {
		go func(taskID int) {
			defer wg.Done()
			time.Sleep(5 * time.Second)
		}(id)
	}

	// Attendre que toutes les tâches soient terminées
	wg.Wait()

	c.JSON(http.StatusOK, gin.H{
		"message": "Toutes les tâches ont été traitées",
		"tasks":   taskIDs,
	})
}
