package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestSetupRouter vérifie que le routeur est bien initialisé
func TestSetupRouter(t *testing.T) {
	router := SetupRouter()
	if router == nil {
		t.Fatal("Le routeur ne doit pas être nil")
	}
}

// TestGetTasksRoute vérifie que l'endpoint "/tasks" fonctionne
func TestGetTasksRoute(t *testing.T) {
	router := SetupRouter()

	req, _ := http.NewRequest("GET", "/tasks", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("GET /tasks: attendu %d, obtenu %d", http.StatusOK, resp.Code)
	}
}
