package main

import (
	"github.com/gin-gonic/gin"
 	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

var DB *gorm.DB

func conectarDB() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := user + ":" + password + "@tcp(" + host + ":3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db
}

func main() {
	godotenv.Load()
	conectarDB()
	DB.AutoMigrate(&Usuario{})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "I'm alive",
		})
	})

	r.GET("/db-status", func(c *gin.Context) {
		sqlDB, err := DB.DB()
		if err != nil {
			c.JSON(500, gin.H{"status": "error", "message": "Failed to get database connection"})
			return
		}
		if err := sqlDB.Ping(); err != nil {
			c.JSON(500, gin.H{"status": "error", "message": "Database is not reachable"})
			return
		}
		c.JSON(200, gin.H{"status": "success", "message": "Database is reachable"})
	})

	r.POST("/agregarusuario", crearUsuario)

	r.GET("/usuarios/:id", obtenerUsuario)

	r.Run(":6769") // listen and serve on 0.0.0.0:6769
}