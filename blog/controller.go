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

	results, err := c.BlogService.GetBlog()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//ctx.JSON(http.StatusOK, results)
	ctx.HTML(http.StatusOK, "blog-view.html", results)
}

// GetEachBlog render Each Blog
func (c *Controller) GetEachBlog(ctx *gin.Context) {
	result, err := c.BlogService.GetEachBlog(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.HTML(http.StatusOK, "blog-post.html", result)
	//ctx.JSON(http.StatusOK, result)
}

// GetCreateUI render frontend interface for blog create
func (c *Controller) GetCreateUI(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "blog-write.html", nil)
}

// MyBlog render frontend interface for my blog
func (c *Controller) MyBlog(ctx *gin.Context) {

	results, err := c.BlogService.GetMyBlog()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//ctx.JSON(http.StatusOK, results)
	ctx.HTML(http.StatusOK, "blog-view.html", results)
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

	ctx.JSON(http.StatusCreated, gin.H{"id": id.Hex(),
		"redirect":    true,
		"redirectUrl": "/blog/myblog"})
}
