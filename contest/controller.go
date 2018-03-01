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

	if c.ContestService.IsVIDExist(ctstInfo.VID) {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": "A contest with this vid already exist."})
		return
	}

	id, err := c.ContestService.Create(ctstInfo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id.Hex()})
}
