package main

import (
	"github.com/gin-gonic/gin"
	"github.com/usuario/talentpitch_api/config"
	"github.com/usuario/talentpitch_api/controllers"
	"github.com/usuario/talentpitch_api/services"
	"net/http"
	"strconv"
)

func main() {
	// Inicializar la conexión a la base de datos
	config.InitDB()

	// Crear el router de Gin
	r := gin.Default()

    // Ruta de prueba
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

	// Rutas para Users
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

    // Rutas para Challenges
	r.GET("/challenges", controllers.GetChallenges)
	r.GET("/challenges/:id", controllers.GetChallenge)
	r.POST("/challenges", controllers.CreateChallenge)
	r.PUT("/challenges/:id", controllers.UpdateChallenge)
	r.DELETE("/challenges/:id", controllers.DeleteChallenge)

    // Rutas para Videos
    r.GET("/videos", controllers.GetVideos)
    r.GET("/videos/:id", controllers.GetVideo)
    r.POST("/videos", controllers.CreateVideo)
    r.PUT("/videos/:id", controllers.UpdateVideo)
    r.DELETE("/videos/:id", controllers.DeleteVideo)

    // Ruta para consumir GPT y llenar la base de datos con un desafío
	r.POST("/gpt/fill-challenges", func(c *gin.Context) {
		count := 1 // Fijamos el número de desafíos a 1 para esta ruta
		context := "" // Puedes pasar un contexto vacío
		err := services.InsertGPTData("challenge", count, context)  // Llamar a la función con los 3 argumentos
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Desafío generado y agregado a la base de datos"})
	})

    // Ruta para generar challenges con GPT
    r.POST("/gpt/generate-challenges/:count", func(c *gin.Context) {
        count, _ := strconv.Atoi(c.Param("count"))   // Convertir la variable de la URL a entero
        context := c.DefaultQuery("context", "")     // Contexto adicional opcional
        err := services.InsertGPTData("challenge", count, context)  // Llamar a la función con los 3 argumentos
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Challenges generados y agregados a la base de datos"})
    })

    // Ruta para generar usuarios con GPT
    r.POST("/gpt/generate-users/:count", func(c *gin.Context) {
        count, _ := strconv.Atoi(c.Param("count"))   // Convertir la variable de la URL a entero
        context := c.DefaultQuery("context", "")     // Contexto adicional opcional
        err := services.InsertGPTData("user", count, context)  // Llamar a la función con los 3 argumentos
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Usuarios generados y agregados a la base de datos"})
    })

    // Ruta para generar videos con GPT
    r.POST("/gpt/generate-videos/:count", func(c *gin.Context) {
        count, _ := strconv.Atoi(c.Param("count"))   // Convertir la variable de la URL a entero
        context := c.DefaultQuery("context", "")     // Contexto adicional opcional
        err := services.InsertGPTData("video", count, context)  // Llamar a la función con los 3 argumentos
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Videos generados y agregados a la base de datos"})
    })

	// Escuchar en el puerto 8080
	r.Run(":8080")
}
