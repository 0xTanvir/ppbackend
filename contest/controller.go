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
// @Router /contest [post]
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

// Get a contest by id
// @Router /contest/{id} [get]
func (c *Controller) Get(ctx *gin.Context) {
	result, err := c.ContestService.Get(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// Update updates a contest by id
// @Router /contest/{id} [put]
func (c *Controller) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var ctst Ctst
	err := ctx.Bind(&ctst)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !c.ContestService.IsExist(id) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "contest not found"})
		return
	}

	err = c.ContestService.Update(id, &ctst)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "contest updated")
}

// Delete deletes an contest by id
// @Router /organization/{id} [delete]
func (c *Controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if !c.ContestService.IsExist(id) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "contest not found"})
		return
	}

	err := c.ContestService.Remove(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "contest deleted")
}

// Find is a organization listing service with search/filter options
// @Router /contest [get]
func (c *Controller) Find(ctx *gin.Context) {
	var qs QueryFilter

	err := ctx.Bind(&qs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var result []*Ctst
	result, err = c.ContestService.Find(qs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// GetUpcomingContest gather all upcoming contest rom db
// @Router /contest [get]
func (c *Controller) GetUpcomingContest(ctx *gin.Context) {
	results, err := c.ContestService.GetUpcomingContest()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, results)
}
