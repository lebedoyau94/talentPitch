package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/usuario/talentpitch_api/models"
	"github.com/usuario/talentpitch_api/config"
)

// Prueba para crear un desafío exitosamente
func TestCreateChallengeSuccess(t *testing.T) {
	config.InitTestDB()

	r := SetupRouter()
	r.POST("/challenges", CreateChallenge)

	challenge := models.Challenge{
		Title:       "Desafío 1",
		Description: "Este es un desafío para pruebas",
		Difficulty:  3,
		UserID:      1,
	}

	jsonValue, _ := json.Marshal(challenge)
	req, _ := http.NewRequest("POST", "/challenges", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	assert.Equal(t, "Desafío 1", response["title"])
}

// Prueba para crear un desafío con datos inválidos
func TestCreateChallengeInvalidData(t *testing.T) {
	config.InitTestDB()

	r := SetupRouter()
	r.POST("/challenges", CreateChallenge)

	// Desafío sin título
	challenge := models.Challenge{
		Description: "Este es un desafío inválido",
		Difficulty:  5,
	}

	jsonValue, _ := json.Marshal(challenge)
	req, _ := http.NewRequest("POST", "/challenges", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code) // Se espera un error por falta de título
}

// Prueba para obtener un desafío exitosamente
func TestGetChallengeSuccess(t *testing.T) {
	config.InitTestDB()

	r := SetupRouter()
	r.GET("/challenges/:id", GetChallenge)

	// Simular la creación de un desafío
	challenge := models.Challenge{
		Title:       "Desafío 2",
		Description: "Descripción del desafío 2",
		Difficulty:  4,
		UserID:      1,
	}
	config.DB.Create(&challenge)

	req, _ := http.NewRequest("GET", "/challenges/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	assert.Equal(t, "Desafío 2", response["title"])
}

// Prueba para obtener un desafío inexistente
func TestGetChallengeNotFound(t *testing.T) {
	config.InitTestDB()

	r := SetupRouter()
	r.GET("/challenges/:id", GetChallenge)

	req, _ := http.NewRequest("GET", "/challenges/999", nil) // Desafío inexistente
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code) // Se espera un 404
}
