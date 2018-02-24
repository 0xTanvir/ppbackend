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
	viper.SetEnvPrefix("up")

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
}
