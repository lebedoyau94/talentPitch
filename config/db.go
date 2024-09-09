package config

import (
    "fmt"
    "log"
    "os"
    "time"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/usuario/talentpitch_api/models"  // Importar los modelos
)

var DB *gorm.DB

// InitDB inicializa la conexión a la base de datos y ejecuta las migraciones
func InitDB() {
    // Definir el Data Source Name (DSN) para MySQL
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    var err error

    // Intentar conectar a la base de datos hasta 5 veces, con un retraso de 5 segundos entre intentos
    for i := 0; i < 5; i++ {
        DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
        if err == nil {
            break
        }
        log.Println("Intentando conectar a la base de datos...")
        time.Sleep(5 * time.Second)
    }

    // Si no se pudo conectar después de 5 intentos, lanzar error
    if err != nil {
        log.Fatalf("Error al conectar a la base de datos: %v", err)
    }

    log.Println("Conexión a la base de datos establecida correctamente.")

    // Verificar si las tablas existen
    if !tableExists("users") || !tableExists("challenges") || !tableExists("videos") {
        log.Println("Ejecutando migraciones, algunas tablas no existen...")

        err = DB.AutoMigrate(&models.User{}, &models.Challenge{}, &models.Video{})
        if err != nil {
            log.Fatalf("Error al migrar los modelos: %v", err)
        }
        log.Println("Migraciones completadas correctamente.")
    } else {
        log.Println("Migraciones omitidas, las tablas ya existen.")
    }

}

// tableExists verifica si una tabla existe en la base de datos
func tableExists(tableName string) bool {
    var result int
    query := fmt.Sprintf("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s'", os.Getenv("DB_NAME"), tableName)
    DB.Raw(query).Scan(&result)

    return result > 0
}