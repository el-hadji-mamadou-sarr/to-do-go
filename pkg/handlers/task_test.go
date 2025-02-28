package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"to-do-go/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Reset de la liste des tâches avant chaque test
func setupTest() {
	tasks = []models.Task{}
}

// Test de la récupération des tâches (GET /tasks)
func TestGetTasksHandler(t *testing.T) {
	setupTest()

	// Création d'un contexte de test
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Appel du handler
	GetTasks(c)

	// Vérification de la réponse
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, "[]", w.Body.String()) // Doit retourner une liste vide au départ
}

// Test de la création d'une tâche (POST /tasks)
func TestCreateTaskHandler(t *testing.T) {
	setupTest()

	// Données de la requête
	task := map[string]string{"title": "Nouvelle tâche", "status": "pending"}
	jsonValue, _ := json.Marshal(task)

	// Création d'une requête HTTP simulée
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Création du contexte Gin de test
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Exécution du handler
	CreateTask(c)

	// Vérifications
	assert.Equal(t, http.StatusCreated, w.Code)

	// Vérification du format JSON retourné
	expectedResponse := `{"id":1,"title":"Nouvelle tâche","status":"pending"}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}

// TestCreateTaskHandlerInvalidJSON vérifie que l'API renvoie une erreur si le JSON est invalide
func TestCreateTaskHandlerInvalidJSON(t *testing.T) {
	setupTest()

	// Requête avec un JSON mal formé
	invalidJSON := `{"title": "Tâche incorrecte", "status": }`
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer([]byte(invalidJSON)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	CreateTask(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// test pour vérifier que l'on peut pas modifier une tache existante
func TestUpdateExistingTask(t *testing.T) {
	setupTest()
	// Création d'une tâche existante
	tasks = append(tasks, models.Task{ID: 1, Title: "Tâche existante", Status: "pending"})

	// Données de la requête
	task := map[string]string{"title": "Tâche modifiée", "status": "completed"}
	jsonValue, _ := json.Marshal(task)

	// Création d'une requête HTTP simulée
	req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Création du contexte Gin de test
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	UpdateTask(c)

	// Vérifications
	assert.Equal(t, http.StatusOK, w.Code)
	expectedResponse := `{"id":1,"title":"Tâche modifiée","status":"completed"}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}
