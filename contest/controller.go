package contest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller handlers of all contest endPoints
type Controller struct {
	ContestService *Service
}

// New creates a new Contest
func (c *Controller) New(ctx *gin.Context) {

	var ctstInfo CtstInfo
	err := ctx.Bind(&ctstInfo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cr, err := c.ContestService.Create(ctstInfo.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, cr)
}
