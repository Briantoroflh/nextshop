package authentication

import (
	"net/http"
	"nextshop/cmd/database"
	"nextshop/entities"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	var user entities.Users
	if err := database.GetDB().Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not found"})
		return
	}

	if err := user.CheckPassword(input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login Success",
		"data" : &user,
	})

}