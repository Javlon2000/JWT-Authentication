package main

import (
	"github.com/Javlon2000/JWT-Authentication/signup"
	"github.com/Javlon2000/JWT-Authentication/models"
	"github.com/Javlon2000/JWT-Authentication/login"
	"github.com/Javlon2000/JWT-Authentication/change"
	"github.com/Javlon2000/JWT-Authentication/websocket"
	
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	models.ConnectDatabase()

	router.POST("/signup", signup.SignUp)

	router.POST("/login", login.Login)

	router.POST("/changepassword", changepassword.ChangePassword)

	router.GET("/echo", websocket.Echo)

	router.Run()
}

