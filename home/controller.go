package home

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Controller is a Homeservice controller
type Controller struct {
	HomeService *Service
}

// GetHomeUI just an example
func (c *Controller) GetHomeUI(ctx *gin.Context) {

	//msg := c.HomeService.GetUI()
	ctx.HTML(http.StatusOK,"index.html",nil)
}

// GetRegistationUI just an example
func (c *Controller) GetRegistationUI(ctx *gin.Context) {
	ctx.HTML(http.StatusOK,"login.html",nil)
}
