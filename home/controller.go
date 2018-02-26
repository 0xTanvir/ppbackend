package home

import "github.com/gin-gonic/gin"

type Controller struct {
	HomeService *Service
}

func (c *Controller) New(ctx *gin.Context) {

	msg := c.HomeService.GetUI()

	ctx.JSON(200, gin.H{
		"message": msg,
	})
}
