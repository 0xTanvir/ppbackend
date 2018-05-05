package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller handlers of all contest endPoints
type Controller struct {
	BlogService *Service
}

// GetUI render frontend interface
func (c *Controller) GetUI(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "blog.html", nil)
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

	id, err := c.BlogService.Create(post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id.Hex()})
}
