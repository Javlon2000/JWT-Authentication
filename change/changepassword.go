package changepassword

import (
	"time"
	"net/http"

	"github.com/Javlon2000/JWT-Authentication/models"
	s "github.com/Javlon2000/JWT-Authentication/signup"

	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
)

type ChangePasswordInput struct {
	Email string `json:"email" binding:"required"`
	CurrentPassword string `json:"currentPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type Payload struct {
	Email string
	Role int
	jwt.StandardClaims
}

var jwtKey = []byte("access_token")

func ChangePassword(c *gin.Context) {

	
	var input ChangePasswordInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
	}
	
    row := models.DB.Table("users").Where("email = ? AND password = ?", input.Email, input.CurrentPassword).Update("password",input.NewPassword)

	if !models.IsNotFound(row) {
    		
		expirationTime := time.Now().Add(168 * time.Hour)

    	payload := Payload {
    		Email: input.Email,
    		Role: 1,
    		StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
    	}

    	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

    	tokenString, err := access_token.SignedString(jwtKey)
    
    	if err != nil {
    		panic(err)
    	}
    	
    	c.JSON(http.StatusOK, gin.H{"access_token": tokenString, "user": s.SignUpInput {Email: input.Email} })

	} else {

    	c.JSON(400, gin.H{"email": input.Email})
	}


}