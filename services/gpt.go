package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"


	"github.com/usuario/talentpitch_api/models"
	"github.com/usuario/talentpitch_api/config"
)

// Estructura para la solicitud a la API de GPT
type GPTRequest struct {
	Prompt string `json:"prompt"`
}

// Estructura para la respuesta de la API de GPT
type GPTResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

// FetchGPTData envía una solicitud a la API de GPT y devuelve la respuesta
func FetchGPTData(prompt string) (string, error) {
	apiURL := os.Getenv("GPT_API_URL")
	apiKey := os.Getenv("GPT_API_KEY")

	if apiURL == "" || apiKey == "" {
		return "", fmt.Errorf("las variables de entorno GPT_API_URL o GPT_API_KEY no están configuradas correctamente")
	}

	// Crear el cuerpo de la solicitud
	requestBody, _ := json.Marshal(GPTRequest{
		Prompt: prompt,
	})

	// Crear la solicitud HTTP
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("error al crear la solicitud: %v", err)
	}

	// Configurar los headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	// Enviar la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error al enviar la solicitud: %v", err)
	}
	defer resp.Body.Close()

	// Procesar la respuesta
	var gptResponse GPTResponse
	err = json.NewDecoder(resp.Body).Decode(&gptResponse)
	if err != nil {
		return "", fmt.Errorf("error al procesar la respuesta: %v", err)
	}

	// Log para imprimir la respuesta completa
	fmt.Printf("Respuesta de GPT: %+v\n", gptResponse)

	// Extraer el texto de la respuesta
	if len(gptResponse.Choices) > 0 && gptResponse.Choices[0].Text != "" {
		return gptResponse.Choices[0].Text, nil
	}

	return "", fmt.Errorf("no se encontró respuesta válida")
}

// InsertGPTData inserta datos generados por GPT en la base de datos
func InsertGPTData(entity string, count int, context string) error {
	// Crear un prompt dinámico según la cantidad de datos a generar
	// Crear un prompt dinámico según la cantidad de datos a generar
	var prompt string

	switch entity {
	case "challenge":
		// Prompt para generar desafíos con estructura
		prompt = fmt.Sprintf(`Genera %d desafíos sobre %s para una plataforma de educación. Cada desafío debe seguir esta estructura:
		1. Título: Un título descriptivo del desafío.
		2. Descripción: Una breve descripción del desafío.
		3. Dificultad: Un número del 1 al 5 que indique la dificultad.`, count, context)

	case "user":
		// Prompt para generar usuarios con estructura
		prompt = fmt.Sprintf(`Genera %d usuarios ficticios para una plataforma de educación. Cada usuario debe seguir esta estructura:
		1. Nombre: Un nombre completo.
		2. Correo electrónico: Un correo electrónico ficticio.
		3. Imagen: Una URL de imagen de perfil del usuario.`, count)

	case "video":
		// Prompt para generar videos con estructura
		prompt = fmt.Sprintf(`Genera %d videos sobre %s para una plataforma de educación. Cada video debe seguir esta estructura:
		1. Título: Un título descriptivo del video.
		2. URL: Un enlace ficticio al video.
		3. Descripción: Una breve descripción del contenido del video.`, count, context)
	}

	// Añadir el contexto si no está vacío
	if context != "" {
		prompt = fmt.Sprintf("%s. Contexto adicional: %s", prompt, context)
	}

	// Obtener la respuesta de GPT y procesar cada resultado
	for i := 0; i < count; i++ {
		data, err := FetchGPTData(prompt)
		if err != nil {
			return fmt.Errorf("error al obtener datos de GPT: %v", err)
		}

		// Insertar los datos en la tabla correspondiente
		switch entity {
		case "challenge":
			challenge := models.Challenge{
				Title:       fmt.Sprintf("Desafío %d generado por GPT", i+1),
				Description: data,
				Difficulty:  3,  // Personalizar según sea necesario
				UserID:      1,  // Se puede ajustar el ID del usuario
			}
			if err := config.DB.Create(&challenge).Error; err != nil {
				return fmt.Errorf("error al insertar los datos en la base de datos: %v", err)
			}
		case "user":
			user := models.User{
				Name:      fmt.Sprintf("Usuario %d generado por GPT", i+1),
				Email:     fmt.Sprintf("user%d@example.com", i+1), // Generar emails únicos
				ImagePath: "/images/default.jpg",
			}
			if err := config.DB.Create(&user).Error; err != nil {
				return fmt.Errorf("error al insertar los datos en la base de datos: %v", err)
			}
		case "video":
			video := models.Video{
				Title:  fmt.Sprintf("Video %d generado por GPT", i+1),
				URL:    fmt.Sprintf("http://example.com/video%d.mp4", i+1), // Generar URLs únicos
				UserID: 1,  // Ajustar el ID de usuario
			}
			if err := config.DB.Create(&video).Error; err != nil {
				return fmt.Errorf("error al insertar los datos en la base de datos: %v", err)
			}
		}
	}

	return nil
}
