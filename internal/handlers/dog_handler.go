package handlers

import (
	config "main/configs"

	"main/internal/entities"

	"github.com/gin-gonic/gin"
)

func GetDogs(c *gin.Context) {
	var dogs []entities.Dog
	config.Database.Find(&dogs)

	c.JSON(200, dogs)
}

func GetDog(c *gin.Context) {
	var dog entities.Dog
	config.Database.First(&dog, c.Param("id"))

	c.JSON(200, dog)
}

func CreateDog(c *gin.Context) {
	var dog entities.Dog

	if err := c.ShouldBindJSON(&dog); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	config.Database.Create(&dog)

	c.JSON(200, dog)
}

func UpdateDog(c *gin.Context) {
	var dog entities.Dog

	if err := c.ShouldBindJSON(&dog); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	config.Database.Where("id = ?", c.Param("id")).Updates(&dog)

	c.JSON(200, dog)
}

func DeleteDog(c *gin.Context) {
	var dog entities.Dog
	config.Database.Delete(&dog, c.Param("id"))

	c.JSON(200, dog)
}
