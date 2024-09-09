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

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// Prueba para crear un usuario correctamente (caso de éxito)
func TestCreateUserSuccess(t *testing.T) {
	config.InitTestDB() // Inicializa la base de datos de pruebas

	r := SetupRouter()
	r.POST("/users", CreateUser)

	user := models.User{
		Name:      "Juan Pérez",
		Email:     "juan.perez@example.com",
		ImagePath: "/images/juan.jpg",
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	assert.Equal(t, "Juan Pérez", response["name"])
	assert.Equal(t, "juan.perez@example.com", response["email"])
}

// Prueba para crear un usuario con datos inválidos (caso de error)
func TestCreateUserInvalidData(t *testing.T) {
	config.InitTestDB()

	r := SetupRouter()
	r.POST("/users", CreateUser)

	// Enviar un usuario sin nombre (esto debería causar un error)
	user := models.User{
		Email:     "user@invalid.com",
		ImagePath: "/images/invalid.jpg",
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verificar que el código de respuesta sea 400 (Bad Request)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Prueba para obtener un usuario existente (caso de éxito)
func TestGetUserSuccess(t *testing.T) {
	config.InitTestDB()

	r := SetupRouter()
	r.GET("/users/:id", GetUser)

	// Simular la creación de un usuario
	user := models.User{
		Name:      "Carlos Garcia",
		Email:     "carlos@example.com",
		ImagePath: "/images/carlos.jpg",
	}
	config.DB.Create(&user)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	assert.Equal(t, "Carlos Garcia", response["name"])
	assert.Equal(t, "carlos@example.com", response["email"])
}

// Prueba para obtener un usuario inexistente (caso de error)
func TestGetUserNotFound(t *testing.T) {
	config.InitTestDB()

	r := SetupRouter()
	r.GET("/users/:id", GetUser)

	req, _ := http.NewRequest("GET", "/users/999", nil) // Usuario que no existe
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verificar que el código de respuesta sea 404 (Not Found)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
