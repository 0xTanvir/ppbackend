package server

import (
	"fmt"

	"github.com/0xTanvir/pp/users"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Controllers struct {
	User *users.Controller
}

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

	// API endpoints
	// Account registration (non-authenticated)
	join := s.Engine.Group("/join")
	{
		join.GET("", s.Controllers.User.New)
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
