package cfg

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

func init() {
	// Allow config via environment variables
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("pp")

	// Server settings
	viper.SetDefault("port", 3030)
	viper.SetDefault("path", ".")
	viper.SetDefault("domain", "localhost")

	// CORS
	viper.SetDefault("cors.allow.allorigins", true)
	viper.SetDefault("cors.allow.origins", []string{})
	viper.SetDefault("cors.allow.methods", []string{"GET", "PUT", "POST", "DELETE"})
	viper.SetDefault("cors.allow.headers", []string{"Origin", "Authorization", "Content-Length", "Content-Type", "Accept", "Accept-Language"})
	viper.SetDefault("cors.expose.headers", []string{"Content-Length", "Content-Type"})
	viper.SetDefault("cors.maxage", 12*time.Hour)

	// log settings.
	// level can be ERROR|WARNING|INFO|DEBUG.
	viper.SetDefault("log.level", "ERROR")
	viper.SetDefault("log.ginrus", true)

	// MongoDB settings.
	// URI in format 'mongodb://USER:PASSWg@HOST:PORT,HOST:PORT/DBNAME"'.
	viper.SetDefault("db.uri", "mongodb://localhost/TEST")
	viper.SetDefault("db.tls.enable", false)

	// Virtual judge settings
	viper.SetDefault("judge.host", "vjudge.net")
	viper.SetDefault("judge.username", "just_take_it")
	viper.SetDefault("judge.password", "Just4Test")
}
