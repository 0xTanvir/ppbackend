package users

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	UserService *Service
}

func (c *Controller) New(ctx *gin.Context) {

	msg := c.UserService.GetUI()

	ctx.JSON(200, gin.H{
		"message": msg,
	})
}
