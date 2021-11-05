package signup

import (
	"fmt"
	"net/http"
	"net/smtp"
	
	"github.com/Javlon2000/JWT-Authentication/models"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

type SignUpInput struct {
	Email string `json:"email" binding:"required"`
}

func SignUp(c *gin.Context) {

	var input SignUpInput

  	err := c.ShouldBindJSON(&input) 

    getUser := SignUpInput{}

  	if err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
    }

	myuuid := uuid.NewV4()

	Sender := "goguruh01@gmail.com"
	Password := "Qwertyu!op"

	receivers := []string {
		input.Email,
	}

	host := "smtp.gmail.com"
	port := "587"

	auth := smtp.PlainAuth("", Sender, Password, host)

	message := []byte(myuuid.String())

	err = smtp.SendMail(host + ":" + port, auth, Sender, receivers, message)

	if err != nil {
		panic(err)
	}

	fmt.Println("Sended!")

    user := models.User{Email: input.Email, Password: myuuid.String()}
  	
  	row := models.DB.Table("users").Select("email").Where("email = ?", input.Email).Find(&getUser)

  	if !models.IsNotFound(row) {
		
		c.JSON(http.StatusBadRequest, "email already exits")
    	return
	
	} else {

  		models.DB.Table("users").Create(&user)
		c.JSON(http.StatusOK, "Created!")
	}
}