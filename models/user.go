package models

import "gorm.io/gorm"

// User representa el modelo de usuario en la base de datos
type User struct {
    gorm.Model
    Name      string `json:"name"`
    Email     string `json:"email" gorm:"unique"`
    ImagePath string `json:"image_path"`
}