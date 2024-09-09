package models

import "gorm.io/gorm"

// Video representa el modelo de un video en la base de datos
type Video struct {
	gorm.Model
	Title       string `json:"title"`
	URL         string `json:"url"`
	UserID      uint   `json:"user_id"`
	ChallengeID *uint  `json:"challenge_id"` // Es opcional si no está relacionado con un desafío
}
