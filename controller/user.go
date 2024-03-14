package controller

import (
	"github.com/Nezent/Go-PSQL/config"
	"github.com/Nezent/Go-PSQL/models"
	"github.com/gin-gonic/gin"
)

// Fetching User Data
func GetUser(c *gin.Context){
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}


// Creating New User
func CreateUser(c *gin.Context){
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200,&user)
}


// Deleting User based on ID
func DeleteUser(c *gin.Context){
	var user models.User
	c.BindJSON(&user)
	config.DB.Where("ID = ?", c.Param("id")).Delete(&user)
	c.JSON(200,&user)
}


// Updating User based on ID
func UpdateUser(c *gin.Context){
	var user models.User
	config.DB.Where("ID = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200,&user)	
}