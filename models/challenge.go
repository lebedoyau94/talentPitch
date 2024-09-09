package models

import "gorm.io/gorm"

// Challenge representa el modelo de un desafío en la base de datos
type Challenge struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Difficulty  int    `json:"difficulty"`
	UserID      uint   `json:"user_id"`
}
