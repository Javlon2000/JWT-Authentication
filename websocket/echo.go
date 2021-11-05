package websocket

import (
	"net/http"
	"gopkg.in/olahol/melody.v1"

	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
)

type Payload struct {
	Email string
	Role int
	jwt.StandardClaims
}

var jwtKey = []byte("access_token")

func Echo(c *gin.Context) {

	webSocketRouter := melody.New()

	token := c.GetHeader("accesToken")
	
	payload := &Payload{}

	// webSocketRouter.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	tkn, err := jwt.ParseWithClaims(token, payload, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil	
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, "StatusUnauthorized")
			return
		}
		c.JSON(http.StatusBadRequest, "StatusBadRequest")
		return
	}
	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, "StatusUnauthorized")
		return
	}

	c.JSON(200, "Verified!")

	webSocketRouter.HandleMessage(func (s *melody.Session, m []byte){
		webSocketRouter.Broadcast(m)
	})
}