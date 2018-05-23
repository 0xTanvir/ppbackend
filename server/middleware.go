package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/0xTanvir/pp/auth"
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

// Middleware collection struct.
type Middleware struct {
	Auth *auth.Service
}



// ReqAuthUser returns middleware which requires authenticated user for request.
func (m *Middleware) ReqAuthUser(c *gin.Context) {

	tokenString, err := c.Cookie("access-token")
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["HS256"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("secure-development-key"), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid{
		c.AbortWithError(401, err)
		return
	}

	// Store claims in auth service
	m.Auth.SetUserID(claims["user"].(string))
	m.Auth.SetName(claims["name"].(string))

	c.Next()
}