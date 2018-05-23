package home

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

// Controller is a Homeservice controller
type Controller struct {
	HomeService *Service
}

func (c *Controller) LoggedIn(ctx *gin.Context) bool {
	tokenString, err := ctx.Cookie("access-token")
	if err != nil {
		return false
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["HS256"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("secure-development-key"), nil
	})
	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid{
		return false
	}
	return true
}

// GetHomeUI just an example
func (c *Controller) GetHomeUI(ctx *gin.Context) {
	if c.LoggedIn(ctx){
		// Render Authorized index
		ctx.HTML(http.StatusOK, "index.html", "nav_auth")
	}else {
		// Render Unauthorized index
		ctx.HTML(http.StatusOK, "index.html", "nav_public")
	}
	//msg := c.HomeService.GetUI()
}

// GetRegistationUI just an example
func (c *Controller) GetRegistationUI(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}
