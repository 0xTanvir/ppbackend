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

	// Before create user there will not be any mongodb ID,
	// So for create we will check with email to verify user availability
	if c.UserService.EmailExist(userInfo.Email) {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": "Email is already registered"})
		return
	}

	id, err := c.UserService.Create(&userInfo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id.Hex(),
		"redirect":    true,
		"redirectUrl": "/"})
}

// Login User
func (c *Controller) Login(ctx *gin.Context) {

	var login Login

	err := ctx.Bind(&login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Binding"})
		return
	}

	ssion, err := c.UserService.Login(&login)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// set cookie
	ctx.SetCookie("access-token", ssion.Token, 21600,"","",false,false)

	ctx.JSON(http.StatusOK, gin.H{
		"token":   ssion.Token,
		"expiry":  ssion.Expiry,
		"ID":    ssion.ID,
		"redirect":    true,
		"redirectUrl": "/",
	})
}