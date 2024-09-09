package controllers

import (
	"net/http"
	"strconv" // Importa strconv
	"github.com/gin-gonic/gin"
	"github.com/usuario/talentpitch_api/config"
	"github.com/usuario/talentpitch_api/models"
)

// Obtener todos los videos (con paginación)
func GetVideos(c *gin.Context) {
	var videos []models.Video
	limit := 10
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))  // Usar strconv.Atoi
	offset := (page - 1) * limit

	if err := config.DB.Limit(limit).Offset(offset).Find(&videos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, videos)
}

// Obtener un video por ID
func GetVideo(c *gin.Context) {
	var video models.Video
	id := c.Param("id")

	if err := config.DB.First(&video, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video no encontrado"})
		return
	}
	c.JSON(http.StatusOK, video)
}

// Crear un nuevo video
func CreateVideo(c *gin.Context) {
	var video models.Video

	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, video)
}

// Actualizar un video
func UpdateVideo(c *gin.Context) {
	var video models.Video
	id := c.Param("id")

	if err := config.DB.First(&video, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video no encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, video)
}

// Eliminar un video
func DeleteVideo(c *gin.Context) {
	var video models.Video
	id := c.Param("id")

	if err := config.DB.First(&video, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video no encontrado"})
		return
	}

	if err := config.DB.Delete(&video).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Video eliminado"})
}
