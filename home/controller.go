package home

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Controller is a Homeservice controller
type Controller struct {
	HomeService *Service
}

// New just an example
func (c *Controller) New(ctx *gin.Context) {

	//msg := c.HomeService.GetUI()
	ctx.HTML(http.StatusOK,"index.html",nil)
}
