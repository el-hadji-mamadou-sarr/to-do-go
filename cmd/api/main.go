package main

import (
	"net/http"
	"to-do-go/pkg/models"

	"github.com/gin-gonic/gin"
)
// Modèle de données pour une tâche
type Task struct {
	ID          string `json:"id"`        
	Title       string `json:"title"`       
	Description string `json:"description"` 
	Completed   bool   `json:"completed"`    
}

func main() {
	// Initialiser le routeur Gin
	router := gin.Default()

	// Liste de tâches en mémoire
	tasks := []Task{
		{ID: "1", Title: "Apprendre Go", Description: "Commencer par les bases de Go", Completed: false},
		{ID: "2", Title: "Créer une API REST", Description: "Implémenter une API REST avec Gin", Completed: true},
		{ID: "3", Title: "Tester l'API", Description: "Tester les endpoints avec curl ou Postman", Completed: false},
	}

	// Définir l'endpoint GET /tasks
	router.GET("/tasks", func(c *gin.Context) {
		// Retourner la liste des tâches au format JSON
		c.JSON(http.StatusOK, gin.H{
			"tasks": tasks,
		})
	})

	// Démarrer le serveur sur le port 8080
	router.Run(":8080")
}