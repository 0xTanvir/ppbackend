package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller is UserService controller
type Controller struct {
	UserService *Service
}

// New User Creation
func (c *Controller) New(ctx *gin.Context) {

	var userInfo UserInfo

	err := ctx.Bind(&userInfo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Binding"})
		return
	}

	id, err := c.UserService.Create(&userInfo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Creating Problem"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id.Hex(),
		"redirect":    true,
		"redirectUrl": "/"})
}
