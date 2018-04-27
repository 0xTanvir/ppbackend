package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Server handles API requests.
type Server struct {
	Engine      *gin.Engine
	Controllers *Controllers
}

// Run server
func (s *Server) Run() error {
	// Cross-Origin Resource Sharing (CORS) Middleware
	cc := cors.Config{
		AllowMethods:  viper.GetStringSlice("cors.allow.methods"),
		AllowHeaders:  viper.GetStringSlice("cors.allow.headers"),
		ExposeHeaders: viper.GetStringSlice("cors.expose.headers"),
		MaxAge:        viper.GetDuration("cors.maxage"),
	}

	// When developing you might want to allow all origins, in production this
	// option must not be used!
	if viper.GetBool("cors.allow.allorigins") {
		cc.AllowAllOrigins = true
	} else {
		cc.AllowOrigins = viper.GetStringSlice("cors.allow.origins")
	}

	// Add CORS Middleware to server's engine
	s.Engine.Use(cors.New(cc))

	// Install Version Middleware
	// s.Engine.Use(s.IncludeVersion())

	// ============================================================
	// API endpoints

	s.Engine.GET("", s.Controllers.Home.GetHomeUI) //Render ui for user

	// -----------Non-authenticated route------------

	// join is used for sign up or regestation
	join := s.Engine.Group("/join")
	{
		join.GET("", s.Controllers.Home.GetRegistationUI) //Render ui for user
		join.POST("", s.Controllers.User.New)
	}

	// login is used for sign in
	login := s.Engine.Group("/login")
	{
		login.GET("")
		login.POST("")
	}

	auth := s.Engine.Group("/auth")
	{
		auth.GET("/refresh")

		// Password management group
		password := auth.Group("/password")
		{
			password.GET("/reset") //Render ui for user
			password.POST("/reset")
			password.POST("/set")
		}
	}

	// Authenticated route
	//TODO this route should be used a ReqAuthUser middleware
	account := s.Engine.Group("/account")
	{
		account.GET("") //Render ui for user
		account.POST("")
		account.GET("/:id")
		account.PUT("/:id")
		account.DELETE("/:id")
	}

	contest := s.Engine.Group("/contest")
	{
		contest.POST("", s.Controllers.Contest.New)
		contest.GET("/:id", s.Controllers.Contest.Get)
		contest.PUT("/:id", s.Controllers.Contest.Update)
		contest.DELETE("/:id", s.Controllers.Contest.Delete)
	}

	return s.Engine.Run(fmt.Sprintf("%v:%v", viper.GetString("host"), viper.GetString("port")))
}

/*
// IncludeVersion will auto set version information header.
func (s *Server) IncludeVersion(args ...string) gin.HandlerFunc {
	// Set out header value for each response
	return func(c *gin.Context) {
		if len(s.Version) > 0 {
			c.Writer.Header().Set("X-Version", s.Version)
		}
		if len(s.BuildDate) > 0 {
			c.Writer.Header().Set("X-BuildDate", s.BuildDate)
		}
		c.Next()
	}
}*/
