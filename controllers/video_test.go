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

// Prueba para crear un video exitosamente
func TestCreateVideoSuccess(t *testing.T) {
	config.InitTestDB()

	r := SetupRouter()
	r.POST("/videos", CreateVideo)

	video := models.Video{
		Title:       "Video 1",
		URL:         "http://example.com/video1.mp4",
		UserID:      1,
		ChallengeID: nil,
	}

	jsonValue, _ := json.Marshal(video)
	req, _ := http.NewRequest("POST", "/videos", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	assert.Equal(t, "Video 1", response["title"])
}

// Prueba para obtener un video exitosamente
func TestGetVideoSuccess(t *testing.T) {
	config.InitTestDB()

	r := SetupRouter()
	r.GET("/videos/:id", GetVideo)

	// Simular la creación de un video
	video := models.Video{
		Title:  "Video 2",
		URL:    "http://example.com/video2.mp4",
		UserID: 1,
	}
	config.DB.Create(&video)

	req, _ := http.NewRequest("GET", "/videos/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	assert.Equal(t, "Video 2", response["title"])
}
