package main

import (
	"github.com/gin-gonic/gin"
 	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Nombre  string `json:"nombre" gorm:"not null"`
	Email   string `json:"email" gorm:"unique;not null"`
	Edad    int    `json:"edad" gorm:"not null"`
}

func crearUsuario(c *gin.Context) {
	var nuevoUsuario Usuario
	if err := c.BindJSON(&nuevoUsuario); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := DB.Create(&nuevoUsuario).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(201, nuevoUsuario)
}

func obtenerUsuario(c *gin.Context) {
	var usuario Usuario
	id := c.Param("id")
	if err := DB.First(&usuario, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, usuario)
}