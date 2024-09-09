package controllers

import (
	"net/http"
	"strconv" // Importa strconv
	"github.com/gin-gonic/gin"
	"github.com/usuario/talentpitch_api/models"
	"github.com/usuario/talentpitch_api/config"
)

// Obtener todos los challenges (con paginación)
func GetChallenges(c *gin.Context) {
	var challenges []models.Challenge
	limit := 10
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))  // Usar strconv.Atoi
	offset := (page - 1) * limit

	if err := config.DB.Limit(limit).Offset(offset).Find(&challenges).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, challenges)
}

// Obtener un challenge por ID
func GetChallenge(c *gin.Context) {
	var challenge models.Challenge
	id := c.Param("id")

	if err := config.DB.First(&challenge, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Challenge no encontrado"})
		return
	}
	c.JSON(http.StatusOK, challenge)
}

// Crear un nuevo challenge
func CreateChallenge(c *gin.Context) {
	var challenge models.Challenge

	if err := c.ShouldBindJSON(&challenge); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&challenge).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, challenge)
}

// Actualizar un challenge
func UpdateChallenge(c *gin.Context) {
	var challenge models.Challenge
	id := c.Param("id")

	if err := config.DB.First(&challenge, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Challenge no encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&challenge); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&challenge).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, challenge)
}

// Eliminar un challenge
func DeleteChallenge(c *gin.Context) {
	var challenge models.Challenge
	id := c.Param("id")

	if err := config.DB.First(&challenge, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Challenge no encontrado"})
		return
	}

	if err := config.DB.Delete(&challenge).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Challenge eliminado"})
}
