package config

import (
    "fmt"
    "log"
    "os"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var TestDB *gorm.DB

// InitTestDB inicializa la conexión a la base de datos MySQL para pruebas
func InitTestDB() {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        "test_db", // Nombre de la base de datos para pruebas
    )

    var err error
    TestDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error al conectar a la base de datos de pruebas: %v", err)
    }

    log.Println("Conexión a la base de datos de pruebas establecida correctamente.")
}
