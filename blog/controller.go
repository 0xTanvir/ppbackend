package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller handlers of all contest endPoints
type Controller struct {
	PostService *Service
}

// New creates a new Post
// @Router /contest [post]
func (c *Controller) New(ctx *gin.Context) {
	var post Post
	err := ctx.Bind(&post)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := c.PostService.Create(post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id.Hex()})
}
