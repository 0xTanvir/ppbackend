package cmd

import (
	"os"
	"time"

	"github.com/0xTanvir/pp/contest"
	"github.com/0xTanvir/pp/db"
	"github.com/0xTanvir/pp/home"
	"github.com/0xTanvir/pp/server"
	"github.com/0xTanvir/pp/users"

	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run server",
	Long:  "run star the apps, it start the server, this is server entry point",
	Run: func(cmd *cobra.Command, args []string) {
		logger := logrus.StandardLogger()
		if lvl, err := logrus.ParseLevel(viper.GetString("log.level")); err == nil {
			logger.Level = lvl
		}
		logger.Out = os.Stderr

		logrus.Info("main : Started : Initialize Mongo")
		// Start MongoDB
		dbc, err := db.Dial(viper.GetString("db.uri"))
		if err != nil {
			logger.Error(err)
			return
		}
		defer dbc.Close()

		// Creates a gin router with default middleware:
		engine := gin.Default()
		if viper.GetBool("log.ginrus") {
			engine.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))
		}

		// Initialize all available controller with their service
		controllers := &server.Controllers{
			User:    &users.Controller{UserService: &users.Service{DB: dbc}},
			Home:    &home.Controller{HomeService: &home.Service{DB: dbc}},
			Contest: &contest.Controller{ContestService: &contest.Service{DB: dbc}},
		}

		server := &server.Server{
			Engine:      engine,
			Controllers: controllers,
		}

		if err := server.Run(); err != nil {
			logger.Error(err)
			return
		}
	},
}
